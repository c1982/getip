package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

var (
	providers []string
	publicIP  string
	err       error
)

func init() {
	providers = []string{
		"http://checkip.amazonaws.com/",
		"https://api.ipify.org/",
		"http://bot.whatismyipaddress.com/",
		"https://wgetip.com",
		"https://icanhazip.com",
		"https://ipinfo.io/ip",
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			getWith3rdParty()
		}
	}()

	out, _ := exec.Command("nslookup", "-type=TXT", "o-o.myaddr.l.google.com", "ns1.google.com").Output()
	r, _ := regexp.Compile(`[0-9]+(?:\.[0-9]+){3}`)
	print(r.FindAllString(string(out), -1)[1])

}

func getWith3rdParty() {
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
