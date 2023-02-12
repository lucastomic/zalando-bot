package proxy

// proxiesFileRoute is the route where all the peoxies are stored
const proxiesFileRoute string = "../../data/proxies.csv"

// Iterator is a proxies iterator.
// Gets a list of available proxies and returns them one by one.
type Iterator struct {
	proxies      []Proxy
	currentProxy int
}

// NewIterator returns a new PeoxieIterator.
// It reads and uses all the proxies from the data/proxies.csv file
func NewIterator() (Iterator, error) {
	proxyFile := file{proxiesFileRoute}
	proxies, err := proxyFile.readProxies()
	if err != nil {
		return Iterator{}, err
	}

	return Iterator{
		proxies:      proxies,
		currentProxy: 0,
	}, err
}

// RefreshProxies change the proxies with a new list of them which works.
// If there is any error, it returns it
func (p *Iterator) RefreshProxies() error {
	proxyFile := file{proxiesFileRoute}
	err := proxyFile.refreshProxies()
	if err != nil {
		return err
	}
	newProxies, err := proxyFile.readProxies()
	if err != nil {
		return err
	}
	p.proxies = newProxies
	p.currentProxy = 0
	return nil
}

// NextProxy returns the next proxy from the available proxys list.
// Proxys are returned in circular way. Every time NextProxy is called returns
// a proxy and sets the next proxy to return.
// Once all the available proxies have been retured, it starts again for the first one.
func (p *Iterator) NextProxy() Proxy {
	var res Proxy
	if !p.hasNext() {
		p.currentProxy = 0
	}
	res = p.proxies[p.currentProxy]
	p.currentProxy++
	return res
}

// hasNext returns if still there are enough proxies to return (the current proxy isn't the last one)
func (p Iterator) hasNext() bool {
	return p.currentProxy < len(p.proxies)
}
