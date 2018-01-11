package main

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/miekg/dns"
)

var dnsServer []string

func init() {
	dnsServer = []string{
		"o-o.myaddr.l.google.com ns1.google.com",
		"myip.opendns.com resolver1.opendns.com",
	}
}

func main() {

	var (
		err      error
		response *dns.Msg
	)

	for _, v := range dnsServer {
		split := strings.Split(v, " ")

		if len(split) != 2 {
			err = errors.New("can't resolve : " + v)
			continue
		}

		local, host := split[0], split[1]

		client, msg := new(dns.Client), new(dns.Msg)

		msg.SetQuestion(dns.Fqdn(local), dns.TypeANY)

		response, _, err = client.Exchange(msg, host+":53")

		if err == nil {
			rx, _ := regexp.Compile(`[0-9]+(?:\.[0-9]+){3}`)

			if len(response.Answer) == 1 {
				print(rx.FindAllString(response.Answer[0].String(), -1)[0])
				break
			}

			err = errors.New("can't resolve")

		}

	}

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

}
