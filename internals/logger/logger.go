package logger

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
	"github.com/go-rod/stealth"
	"github.com/lucastomic/zalando-bot/internals/proxy"
)

const (
	// url              = "https://accounts.zalando.com/authenticate?sales_channel=1e161d6e-0427-4cfc-a357-e2b501188a15&request=eyJjbGllbnRfaWQiOiJmYXNoaW9uLXN0b3JlLXdlYiIsInJlc3BvbnNlX3R5cGUiOiJjb2RlIiwic2NvcGVzIjpbIm9wZW5pZCJdLCJyZWRpcmVjdF91cmkiOiJodHRwczovL3d3dy56YWxhbmRvLmVzL3Nzby9jYWxsYmFjayIsInN0YXRlIjoiZXlKdmNtbG5hVzVoYkY5eVpYRjFaWE4wWDNWeWFTSTZJbWgwZEhCek9pOHZkM2QzTG5waGJHRnVaRzh1WlhNdmJYbGhZMk52ZFc1MEx5SXNJblJ6SWpvaU1qQXlNeTB3TWkwd05WUXhOam95TmpvMU0xb2lmUT09Iiwibm9uY2UiOiI2ZGVkZTI3NC03MGYzLTQ2MWEtYjU4MC05MTMxZTZkNTZiOWIiLCJ1aV9sb2NhbGVzIjpbImVzLUVTIl0sInJlcXVlc3RfaWQiOiI4dnFGZGJMUG8tVExLUE4xOjUwNTNmM2ZmLWNmNmYtNGRjYy1iNGU4LWVlMTdhNDE4M2RmYjp6QjhkQ3ktek5EVkJvUUkyIiwiZiI6WyJQV0NfQSJdfQ==&ui_locales=es-ES&passwordMeterFT=true"
	url              = "https://www.promiedos.com.ar/"
	emailSelector    = "#sso > div > div:nth-child(2) > main > div > div._134xl > div > div > div > form > div:nth-child(1) > div > div > input"
	passwordSelector = "#sso > div > div:nth-child(2) > main > div > div._134xl > div > div > div > form > div:nth-child(2) > div > div > input"
	submitSelector   = "#sso > div > div:nth-child(2) > main > div > div._134xl > div > div > div > form > button"
)

// proxyIterator is the iterator wihch will provide the proxies for the scrapping
var proxyIterator, _ = proxy.NewIterator()

// fillEmail looks for the Email input and fills it with the email var content
func fillEmail(page *rod.Page, email string) {
	page.MustElement(emailSelector).MustInput(email)
}

// fillPassword looks for the Password input and fills it with the password var content
func fillPassword(page *rod.Page, password string) {
	page.MustElement(passwordSelector).MustInput(password)
}

// pressSubmitButton press the submit button to authenticate the user
func pressSubmitButton(page *rod.Page) {
	page.MustElement(submitSelector).MustClick()
}

// getControlURL retunrns the URL of a launcher which uses the proxy with IP passed as argument
func getControlURL(proxy proxy.Proxy) string {
	l := launcher.New()
	l = l.Set(flags.ProxyServer, proxy.GetProxyURL())
	l = l.Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome")
	l.Headless(true)
	controlURL, _ := l.Launch()
	return controlURL
}

// openPage looks for the url and opens it in the navigator.
// Returns the object of type Page.
func openPage() *rod.Page {
	proxy := proxyIterator.NextProxy()
	controlURL := getControlURL(proxy)
	browser := rod.New().ControlURL(controlURL).MustConnect()
	return stealth.MustPage(browser).MustNavigate(url)
}

// Signs in the user with the propierties setted before
func SignIn(email, password string) {
	page := openPage()
	// fillEmail(page, email)
	// fillPassword(page, password)
	// pressSubmitButton(page)

	testF(page)
}

func testF(page *rod.Page) {
	// el := page.MustElement("#z-navicat-header-root > header > div:nth-child(2) > div > div > div > div.GRuH6Q.C3wGFf > div > div > div > div.z-navicat-header_bottomRow > div:nth-child(2) > nav > ul > li:nth-child(2) > span > a > span")
	el := page.MustElement("#titulo1")
	fmt.Print(el.MustText())
}
