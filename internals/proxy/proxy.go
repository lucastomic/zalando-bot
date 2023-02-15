package proxy

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Proxy server
// ip is a stirng with his IP direction. Example: 192.168.1.100
// port is a string with the port wihch is allowed to use. Exaple: 80
// protocols is an string array with the supported protocoles. Example: ["HTTPS", "HTTP"]
type Proxy struct {
	ip        string
	port      string
	protocols []string
}

// newProxy returns a proxy with the data passed as argument
func newProxy(ip string, port string, protocols []string) Proxy {
	return Proxy{ip, port, protocols}
}

// NewProxyFromCSV returns a proxy from a CSV line with the next format:
// Protocol, IP, Port.
// If the string doesn't have this format, it returns an error
// For example, given:
// HTTP, 192.168.1.100, 8080
// it returns:
//
//	Proxy{
//	  ip:192.168.1.100
//	  port:8080
//	  protocols:["HTTP"]
//	}
func newProxyFromCSV(csv string) (Proxy, error) {
	csvSplitted := strings.Split(csv, ",")
	if len(csvSplitted) < 2 {
		return Proxy{}, errors.New("invalid CSV line. It has to have th next format: Protocols,IP,Port")
	}
	return Proxy{
		ip:        csvSplitted[1],
		port:      csvSplitted[2],
		protocols: []string{csvSplitted[0]},
	}, nil
}

// GetProxyURL returns the prozy's URL.
// This means, a string with the next format:
// "proxyIP:proxyPort"
func (p Proxy) GetProxyURL() string {
	return p.ip + ":" + p.port
}

// Parse returns a string with the Proxy info.
// For example: http://103.119.95.106:80
func (p Proxy) Parse() string {
	return fmt.Sprintf("%s://%s:%s", p.protocols[0], p.ip, p.port)
}

// check checks whether the proxy works properly.
// If it doesn't it prints the error or the status response
// To checks this makes a simple request to http://ip-api.com/json/
func (p *Proxy) check() bool {
	resp, err := p.makeSimpleReq()
	if err != nil || resp.StatusCode != 200 {
		return false
	}

	return true
}

// makeSimpleReq makes a simple GET request with the proxy to test if it works properly
func (p Proxy) makeSimpleReq() (*http.Response, error) {
	proxyURL, _ := url.Parse(p.Parse())
	proxy := http.ProxyURL(proxyURL)
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://www.promiedos.com.ar/", nil)

	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

// toCSV returns the a string with the proxy information in CSV fromat with an line jump at the end.
// For example:
// "HTTP, 192.168.1.100, 8080 \n"
func (p Proxy) toCSV() string {
	return p.protocols[0] + "," + p.ip + "," + p.port + "\n"
}
