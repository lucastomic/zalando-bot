package logger

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
	"github.com/lucastomic/zalando-bot/internals/proxy"
)

const (
	url              = "https://accounts.zalando.com/authenticate?sales_channel=1e161d6e-0427-4cfc-a357-e2b501188a15&request=eyJjbGllbnRfaWQiOiJmYXNoaW9uLXN0b3JlLXdlYiIsInJlc3BvbnNlX3R5cGUiOiJjb2RlIiwic2NvcGVzIjpbIm9wZW5pZCJdLCJyZWRpcmVjdF91cmkiOiJodHRwczovL3d3dy56YWxhbmRvLmVzL3Nzby9jYWxsYmFjayIsInN0YXRlIjoiZXlKdmNtbG5hVzVoYkY5eVpYRjFaWE4wWDNWeWFTSTZJbWgwZEhCek9pOHZkM2QzTG5waGJHRnVaRzh1WlhNdmJYbGhZMk52ZFc1MEx5SXNJblJ6SWpvaU1qQXlNeTB3TWkwd05WUXhOam95TmpvMU0xb2lmUT09Iiwibm9uY2UiOiI2ZGVkZTI3NC03MGYzLTQ2MWEtYjU4MC05MTMxZTZkNTZiOWIiLCJ1aV9sb2NhbGVzIjpbImVzLUVTIl0sInJlcXVlc3RfaWQiOiI4dnFGZGJMUG8tVExLUE4xOjUwNTNmM2ZmLWNmNmYtNGRjYy1iNGU4LWVlMTdhNDE4M2RmYjp6QjhkQ3ktek5EVkJvUUkyIiwiZiI6WyJQV0NfQSJdfQ==&ui_locales=es-ES&passwordMeterFT=true"
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
	// l.Headless(true)
	controlURL, _ := l.Launch()
	return controlURL
}

// getBrowser returns the browser which will be used.
// It uses a control URL with a proxy from the proxyIterator
func getBrowser() *rod.Browser {
	var response *rod.Browser
	proxy := proxyIterator.NextProxy()
	controlURL := getControlURL(proxy)
	// var cookies []*proto.NetworkCookieParam = []*proto.NetworkCookieParam{
	// 	{
	// 		Name:  "_abck",
	// 		Value: "DD857431E05E86A63679BB9183A68CC8~-1~YAAQdSkRAtWtg0uGAQAAXWUEVglWbmPci+VrL0l+1JbV48221WdP/8KxivFTGb2D5OZuFIO4VEPds19DuknAdL0fzmgCDF9pSJsfkTN/9+Ph115tDnWDywiqmRKHcGSQVaCdUrJQc5v+DGfRMgzfUHvdArVng+jpL2kD2aIXRvKgTGvUVX9HI1U1XVL7XtzNdeXa7r5TVLNuBSwKWptY1lgOLDLS5eiO8kRRLsnEX9rQrW45/qO4UvjqHapMZD92vnw3cGii70yrV7vcEIlujIbAstOwoULuzCHi7DtW7fiDNkBXPc4LTpA1XEYyOzbRdPEWIsM8BAZRxi+6VGLPyHGH3DwsX1QXIhTC0CW393kT5bX1TX4jOC6E80mtz+iAscoSZ8dYoXK6DsA2TVkEnntuCnWw++yWLAkhjrFm794q26EYqOpF1k7cC0Mw2MutQsN9xdI2nONJbERqvWmUWh0sZgil35V+63tA0u4YE+VE4zAhMAI59d0aS2A2XbKHZOXiQ0qIJLzzAg==~-1~-1~-1",
	// 	},
	// 	{
	// 		Name:  "bm_sz",
	// 		Value: "2BD8AFF9B40CD4DA5DFFF40A8C167EF9~YAAQdSkRAr2Yg0uGAQAAWiwDVhJp4wpdzsdXCt1QC/fkyQOjhqUd3uk/86lC59pldY96tTR3Xxx343TfYG1wW8yzw+F6U2hg8bvQqrYm6SBRaBnrAG+PpyjEwdS5wzELpgw0fdMbYkJKuzV8GWqh32AraCzo5uU5Kol2Lp1LBgL3SX5wPQ9JVj6f4aS+6/Ge83u+Qx7ohzMOs4t+K+Yy1fa83Vr69ZQe3IXV7qgCm7SglPEPhh3gfNar5JxuEvwWGAXff+1woiiJ8GaWyIyiVCS+6s/HyXAatdkw7nDpwF1URUFAkF551FClIKdL0mSJLZyfLPS8WEfMaGCZOWo2uo6ySz8GtxtOa8as8fSO1DDJxjFGZYtVR6nCVM3wU9tU29j/RJkvnw7O4ZfiIq5Co37XvA==~3753526~3552568",
	// 	},
	// }
	response = rod.New().ControlURL(controlURL).MustConnect()
	// response.SetCookies(cookies)
	return response
}

// getPage opens and returns the page that will be used.
// It also sets an User Agent (it is usefull for the akami bypass)
func getPage() *rod.Page {
	browser := getBrowser()
	page := stealth.MustPage(browser).MustNavigate(url)
	page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Linux; U; Android 4.4.2; en-us; SCH-I535 Build/KOT49H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
	})
	return page
}

// Signs in the user with the propierties setted before
func SignIn(email, password string) {
	launcher.NewBrowser().MustGet()
	page := getPage()

	fillEmail(page, email)
	fillPassword(page, password)
	pressSubmitButton(page)

	testF(page)
}

func testF(page *rod.Page) {
	el := page.MustElement("#z-navicat-header-root > header > div:nth-child(2) > div > div > div > div.GRuH6Q.C3wGFf > div > div > div > div.z-navicat-header_bottomRow > div:nth-child(2) > nav > ul > li:nth-child(2) > span > a > span")
	// el := page.MustElement("#titulo1")
	fmt.Println(el.MustText())
}
