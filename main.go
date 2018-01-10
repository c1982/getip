package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

var (
	providers []string
	publicIP  string
	err       error
)

func init() {
	providers = []string{
		"https://api.ipify.org/",
		"http://bot.whatismyipaddress.com/",
		"https://wgetip.com",
		"https://icanhazip.com",
		"https://ipinfo.io/ip",
	}
}

func main() {

	for i := 0; i < len(providers); i++ {
		publicIP, err = get(providers[i])

		if err == nil {
			break
		}
	}

	if err != nil {
		print(err)
		os.Exit(1)
	} else {
		print(publicIP)
	}
}

func get(provider string) (ipaddr string, err error) {

	resp, err := http.Get(provider)

	if err != nil {
		return ipaddr, err
	}

	return toString(resp)
}

func toString(resp *http.Response) (ipaddr string, err error) {

	body, err := ioutil.ReadAll(resp.Body)
	ipaddr = string(body)

	return ipaddr, err
}
