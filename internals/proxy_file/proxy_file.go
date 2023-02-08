package proxyfile

import (
	"os"
	"strings"
	"sync"

	"github.com/lucastomic/zalando-bot/internals/proxy"
)

// ProxyFile is a (CSV) file where proxies are stored.
// Proxies are constantly beeing removed and addded, depending on if they works properly
// or if they stop doing it.
// TODO: They are ordered depending of his latency.

// They are in the next format:
// Protocol, IP, Port

// For example:

// HTTP, 192.168.1.100, 8080
// HTTPS, 10.0.0.1, 8888
// SOCKS, 172.16.0.1, 1080
type ProxyFile struct {
	path string
}

// NewProxyFile returns a ProxyFile with the route passed as argument
func NewProxyFile(route string) ProxyFile {
	return ProxyFile{route}
}

// RefreshProxies gets a list of proxies from an external API,
// and write the ones wihch works properly
// If there is an error, it returns it
func (p ProxyFile) RefreshProxies() error {
	var wg sync.WaitGroup

	proxiesToWrite, err := proxy.GetProxiesFromAPI()
	if err != nil {
		return err
	}

	file, err := os.Create(p.path)
	if err != nil {
		return err
	}
	file.Write([]byte("Protocol, IP, Port \n"))

	for _, proxy := range proxiesToWrite {
		wg.Add(1)
		go p.writeProxieIfWorks(proxy, file, &wg)
	}

	wg.Wait()
	return nil

}

// writeProxieIfWorks writes the proxie passed as argument in proxies.csv if it works properly.
// Also needs a *sync.WaitGroup object to execute it concurrently
func (p ProxyFile) writeProxieIfWorks(proxy proxy.Proxy, file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	if proxy.Check() {
		file.Write([]byte(proxy.ToCSV()))
	}
}

// ReadProxies reads all the proxies from the file and retruns a slice of them.
// If there is an error it returns it as second value.
func (p ProxyFile) ReadProxies() ([]proxy.Proxy, error) {
	//TODO: implement
	var response []proxy.Proxy
	// proxiesFile, err := ioutil.ReadFile(p.path)
	//
	// if err != nil {
	// 	return nil, err
	// }
	//
	// proxiesStringSlice := strings.Split(string(proxiesFile), "\n")
	//
	// for i:=range proxiesStringSlice{
	// 	newProxy := proxy.NewHTTPProxyFromIP(removeCarriageReturn(proxiesStringSlice[i]))
	// 	response = append(response, newProxy)
	// }
	//
	return response, nil
}

// removeCarriageReturn removes the carriage return (\r) from a string
func removeCarriageReturn(s string) string {
	return strings.Replace(s, "\r", "", 1)
}
