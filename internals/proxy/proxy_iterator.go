package proxy

import (
	"fmt"
	"io/ioutil"
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
