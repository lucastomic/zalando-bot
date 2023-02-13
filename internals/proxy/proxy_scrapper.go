package proxy

import (
	"github.com/gocolly/colly"
)

const urlToScrapp string = "https://free-proxy-list.net/"

// getProxiesFromScrapping scrapps the URL specified before and returns the proxies's list
// which are contained in this WEB.
func getProxiesFromScrapping() []Proxy {
	var response []Proxy

	c := colly.NewCollector()
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		response = append(response, getProxyFromRowElement(e))
	})

	c.Visit(urlToScrapp)
	return response
}

// getProxyFromRowElement returns the proxie with the information in the HTML row passed as argument
// The ip is in the first column,
// the port in the second one
// whether it's https or not in the seventh one
func getProxyFromRowElement(row *colly.HTMLElement) Proxy {
	ip := row.ChildText("td:nth-of-type(1)")
	port := row.ChildText("td:nth-of-type(2)")
	isHttps := row.ChildText("td:nth-of-type(7)")

	return newProxy(ip, port, getProtocols(isHttps))
}

// getProtocols returns a slice of strings with the protocols, receiving a string which
// indicates whether or not the protocol is https.
// If it is, the stirng isHttps is "yes", and "no" otherwise
// If is not https is supposed that is http
func getProtocols(isHttps string) []string {
	var protocols []string
	if isHttps == "yes" {
		protocols = append(protocols, "https")
	} else {
		protocols = append(protocols, "http")
	}
	return protocols

}
