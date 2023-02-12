package proxy

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// file is a (CSV) file where proxies are stored.
// Proxies are constantly beeing removed and addded, depending on if they works properly
// or if they stop doing it.
// TODO: They are ordered depending of his latency.
// They are in the next format:
// Protocol, IP, Port
// For example:
// HTTP, 192.168.1.100, 8080
// HTTPS, 10.0.0.1, 8888
// SOCKS, 172.16.0.1, 1080
type file struct {
	path string
}

// RefreshProxies gets a list of proxies from an external API,
// and write the ones wihch works properly
// If there is an error, it returns it
func (p file) refreshProxies() error {
	var wg sync.WaitGroup

	proxiesToWrite, err := getProxiesFromAPI()
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
func (p file) writeProxieIfWorks(proxy Proxy, file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	if proxy.check() {
		file.Write([]byte(proxy.toCSV()))
	}
}

// readProxies reads all the proxies from the file and retruns a slice of them.
// If there is an error it returns it as second value.
func (p file) readProxies() ([]Proxy, error) {
	//TODO: implement
	var response []Proxy
	proxiesFile, err := ioutil.ReadFile(p.path)

	if err != nil {
		return nil, err
	}

	proxiesStringSlice := strings.Split(string(proxiesFile), "\n")
	removeHeader(&proxiesStringSlice)

	for i := range proxiesStringSlice {
		newProxy, err := newProxyFromCSV(proxiesStringSlice[i])
		if err == nil {
			response = append(response, newProxy)
		}
	}

	return response, nil
}

// removeHeader removes the header from the Porxies-CSV-slice, which is the first element.
// This means, as the CSV first line is the header (Protocol, IP, Port), the first element is not a valid proxy,
// but the header, which have to be removed.
func removeHeader(slice *[]string) {
	*slice = (*slice)[1:]
}
