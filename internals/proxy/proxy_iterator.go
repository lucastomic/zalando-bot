package proxy

// Number of available proxies
const proxysNumber = 10

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
