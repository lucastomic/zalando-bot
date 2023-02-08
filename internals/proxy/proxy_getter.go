package proxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var apiURL string = "https://proxylist.geonode.com/api/proxy-list?limit=500&page=1&sort_by=lastChecked&sort_type=desc&google=true"

// GetProxiesFromAPI gets a list of proxies from an external API and returns them in a Proxy slice.
// On error, return null and an error.
func GetProxiesFromAPI() ([]Proxy, error) {
	var proxies []Proxy

	response, err := getProxiesResponse()
	if err != nil {
		return nil, err
	}

	if err := parseProxiesResponse(response, &proxies); err != nil {
		return nil, err
	}

	return proxies, nil
}

// getProxiesResponse makes an HTTP request to get the list of proxies to the API: https://geonode.com/free-proxy-list
// On an error, return null and the error.
// The JSON body from the response looks like this:
//
//	{
//	    "data": [
//	        {
//	            "ip": "110.77.232.221",
//	            "port": "4145",
//	            "protocols": [
//	                "socks4"
//	            ],
//				.
//				.
//				.
//	            "upTimeTryCount": 952
//	        },
//			.
//			.
//			.
//	    ],
//	    "total": 8981,
//	    "page": 1,
//	    "limit": 2
//	}
func getProxiesResponse() (*http.Response, error) {
	res, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// parseProxiesResponse deserializes an HTTP response into a Proxy slice.
// The reponse must have a "data" field which has the proxys objects, this way:

//	{
//	    "data": [
//	        {
//	            "ip": "110.77.232.221",
//	            "port": "4145",
//	            "protocols": [
//	                "socks4"
//	            ],
//	        },
//			.
//			.
//			.
//	    ],
//  }

// On an error, returns it.
func parseProxiesResponse(response *http.Response, proxies *[]Proxy) error {
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var responseData struct {
		Data []struct {
			Ip        string
			Port      string
			Protocols []string
		} `json:"data"`
	}

	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return err
	}

	for _, proxyData := range responseData.Data {
		newPropxy := NewProxy(proxyData.Ip, proxyData.Port, proxyData.Protocols)
		*proxies = append(*proxies, newPropxy)
	}

	return nil
}
