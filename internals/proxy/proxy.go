package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Proxy server. Specifies his IP and his port
type Proxy struct {
	ip        string
	port      string
	protocols []string
}

// NewHTTPProxy returns a HTTP proxy with the ip and the port sepcified as argument
func NewHTTPProxy(host string, port string) Proxy {
	return Proxy{host, port, []string{"HTTP"}}
}

// NewProxy returns a proxy with the data passed as argument
func NewProxy(ip string, port string, protocols []string) Proxy {
	return Proxy{ip, port, protocols}
}

// NewHTTPProxyFromIP returns an HTTP Proxy from a string with the IP host:port.
// For example 103.119.95.106:80
func NewHTTPProxyFromIP(ip string) Proxy {
	splitedIP := strings.Split(ip, ":")
	host := splitedIP[0]
	port := splitedIP[1]

	return NewHTTPProxy(host, port)
}

// getPort is the port getter. Returns a string with the available port of the proxy
func (p Proxy) getPort() string {
	return p.port
}

// getHost is the host getter. Returns a string with the host of the proxy
func (p Proxy) getIP() string {
	return p.ip
}

// Parse returns a string with the Proxy info.
// For example: http:// 103.119.95.106:80
func (p Proxy) Parse() string {
	return fmt.Sprintf("%s://%s:%s", p.protocols[0], p.ip, p.port)
}

// Check checks whether the proxy works properly.
// If it doesn't it prints the error or the status response
// To checks this makes a simple request to http://ip-api.com/json/
func (p *Proxy) Check() bool {
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
	req, _ := http.NewRequest("GET", "http://ip-api.com/json/", nil)
	return client.Do(req)
}

// ToCSV returns the a string with the proxy information in CSV fromat with an line jump at the end.
// For example:
// "HTTP, 192.168.1.100, 8080 \n"
func (p Proxy) ToCSV() string {
	return p.protocols[0] + "," + p.ip + "," + p.port + "\n"
}
