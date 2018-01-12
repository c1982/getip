package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

var (
	resolvers   map[string]string
	defaultPort = 53
	resp        *dns.Msg
	publicIP    string
)

func init() {
	resolvers = map[string]string{
		"o-o.myaddr.l.google.com": "ns1.google.com",
		"myip.opendns.com":        "resolver1.opendns.com",
	}
}

func main() {
	var err error

	client, msg := new(dns.Client), new(dns.Msg)
	msg.RecursionDesired = false

	for name, host := range resolvers {
		msg.SetQuestion(dns.Fqdn(name), dns.TypeANY)
		resp, _, err = client.Exchange(msg, fmt.Sprintf("%s:%d", host, defaultPort))

		if err != nil {
			continue
		}

		if resp.Rcode != dns.RcodeSuccess {
			continue
		}

		publicIP = extractIP(resp)

		if publicIP != "" {
			break
		}
	}

	if err != nil {
		print(err.Error())
		os.Exit(1)
	} else {
		print(publicIP)
	}
}

func extractIP(r *dns.Msg) string {

	for _, rr := range resp.Answer {
		if t, ok := rr.(*dns.TXT); ok {
			return t.Txt[0]
		}

		if a, ok := rr.(*dns.A); ok {
			return a.A.String()
		}
	}

	return ""
}
