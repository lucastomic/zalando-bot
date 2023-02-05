package logger

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
)

const (
	url              = "https://accounts.zalando.com/authenticate?sales_channel=1e161d6e-0427-4cfc-a357-e2b501188a15&request=eyJjbGllbnRfaWQiOiJmYXNoaW9uLXN0b3JlLXdlYiIsInJlc3BvbnNlX3R5cGUiOiJjb2RlIiwic2NvcGVzIjpbIm9wZW5pZCJdLCJyZWRpcmVjdF91cmkiOiJodHRwczovL3d3dy56YWxhbmRvLmVzL3Nzby9jYWxsYmFjayIsInN0YXRlIjoiZXlKdmNtbG5hVzVoYkY5eVpYRjFaWE4wWDNWeWFTSTZJbWgwZEhCek9pOHZkM2QzTG5waGJHRnVaRzh1WlhNdmJYbGhZMk52ZFc1MEx5SXNJblJ6SWpvaU1qQXlNeTB3TWkwd05WUXhOam95TmpvMU0xb2lmUT09Iiwibm9uY2UiOiI2ZGVkZTI3NC03MGYzLTQ2MWEtYjU4MC05MTMxZTZkNTZiOWIiLCJ1aV9sb2NhbGVzIjpbImVzLUVTIl0sInJlcXVlc3RfaWQiOiI4dnFGZGJMUG8tVExLUE4xOjUwNTNmM2ZmLWNmNmYtNGRjYy1iNGU4LWVlMTdhNDE4M2RmYjp6QjhkQ3ktek5EVkJvUUkyIiwiZiI6WyJQV0NfQSJdfQ==&ui_locales=es-ES&passwordMeterFT=true"
	emailSelector    = "#sso > div > div:nth-child(2) > main > div > div._134xl > div > div > div > form > div:nth-child(1) > div > div > input"
	passwordSelector = "#sso > div > div:nth-child(2) > main > div > div._134xl > div > div > div > form > div:nth-child(2) > div > div > input"
	submitSelector   = "#sso > div > div:nth-child(2) > main > div > div._134xl > div > div > div > form > button"
	email            = "lucastomic17@gmail.com"
	password         = "." //Removed to add to Github
)

// fillEmail looks for the Email input and fills it with the email var content
func fillEmail(page *rod.Page) {
	page.MustElement(emailSelector).MustInput(email)
}

// fillPassword looks for the Password input and fills it with the password var content
func fillPassword(page *rod.Page) {
	page.MustElement(passwordSelector).MustInput(password)
}

// pressSubmitButton press the submit button to authenticate the user
func pressSubmitButton(page *rod.Page) {
	page.MustElement(submitSelector).MustClick()
}

// getControlURL retunrns the URL of a launcher which uses the proxy with IP passed as argument
func getControlURL(proxyIP, port string) string {
	l := launcher.New()
	l = l.Set(flags.ProxyServer, proxyIP+":"+port)
	controlURL, _ := l.Launch()
	return controlURL
}

// openPage looks for the url and opens it in the navigator.
// Returns the object of type Page.
func openPage() *rod.Page {
	controlURL := getControlURL("154.236.184.71", "1981")
	return rod.New().MustConnect().ControlURL(controlURL).NoDefaultDevice().MustConnect().MustPage(url)
}

// Signs in the user with the propierties setted before
func SignIn() {
	page := openPage()
	page.MustWindowFullscreen()
	fillEmail(page)
	fillPassword(page)
	pressSubmitButton(page)
	time.Sleep(time.Hour)
}
