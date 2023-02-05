package proxy

// Proxy server. Specifies his IP and his port
type Proxy struct {
	IP   string
	port string
}

// NewProxy retunrs a proxy with the ip and the port sepcified as argument
func NewProxy(ip, port string) Proxy {
	return Proxy{ip, port}
}

// getPort is the port getter. Returns a string with the available port of the proxy
func (p Proxy) getPort() string {
	return p.port
}

// getIP is the IP getter. Returns a string with the IP of the proxy
func (p Proxy) getIP() string {
	return p.IP
}
