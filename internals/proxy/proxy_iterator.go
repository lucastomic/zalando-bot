package proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Number of available proxies
const proxysNumber = 50

// List of available proxies
var proxys [proxysNumber]Proxy

// Index of the proxie will be returned once NextProxy() is called
var currentProxy = 0

// NextProxy returns the next proxy from the available proxys list.
// Proxys are returned in circular way. Every time NextProxy is called returns
// a proxy and sets the next proxy to return.
// Once all the available proxies have been retured, it starts again for the first one.
func NextProxy() Proxy {
	var res Proxy
	if hasNext() {
		res = proxys[currentProxy]
		currentProxy++
	} else {
		currentProxy = 0
		res = proxys[currentProxy]
	}
	return res
}

// hasNext returns if still there are enough proxies to return
func hasNext() bool {
	return currentProxy < proxysNumber
}

// Inits the proxies  with the
func InitProxies() error {
	proxies, err := ioutil.ReadFile("../../assets/http_proxies.txt")
	if err != nil {
		fmt.Println(err)
		return err
	}
	proxiesString := strings.Split(string(proxies), "\n")
	for i := 0; i < proxysNumber; i++ {
		newProxy := NewHTTPProxyFromIP(removeCarriageReturn(proxiesString[i]))
		proxys[i] = newProxy
	}

	return nil
}

// removeCarriageReturn removes the carriage return (\r) from a string
func removeCarriageReturn(s string) string {
	return strings.Replace(s, "\r", "", 1)
}

// FilterProxies filter the proxies which doesn't works
func FilterProxies() {
	proxies, err := ioutil.ReadFile("../../assets/http_proxies.txt")
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create("./proxies.txt")

	proxiesString := strings.Split(string(proxies), "\n")

	for i := 0; i < proxysNumber; i++ {
		newProxy := NewHTTPProxyFromIP(removeCarriageReturn(proxiesString[i]))
		fmt.Println(newProxy.Check())
		if newProxy.Check() {
			f.Write([]byte(newProxy.Parse()))
		}
		proxys[i] = newProxy
		fmt.Println(proxys[i].Parse())
	}

}

// GetProxies gets a list of proxies from an external API and returns them in a Proxy segment.
// On error, return null and an error.
func GetProxies() ([]Proxy, error) {
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
// On an error, return null and an error.
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
	url := fmt.Sprintf("https://proxylist.geonode.com/api/proxy-list?limit=%d&page=1&sort_by=lastChecked&sort_type=desc", proxysNumber)
	res, err := http.Get(url)
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
