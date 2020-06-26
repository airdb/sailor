package sailor

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/miekg/dns"
)

const (
	DefaultDNSServer = "8.8.8.8:53"
)

var DNSTimeout time.Duration = 3
var DNSRetry int = 3

// FQDN: A fully qualified domain name.
func SetFQDN(domain string) string {
	if !strings.HasSuffix(domain, ".") {
		domain += "."
	}

	return domain
}

func QueryDNSSRVRecord(domain string) {
	c := dns.Client{Timeout: DNSTimeout * time.Second}

	m := dns.Msg{}
	m.SetQuestion(SetFQDN(domain), dns.TypeSRV)

	r, _, err := c.Exchange(&m, DefaultDNSServer)
	if err != nil {
		log.Fatalf("query SRV record failed, domain: %s, err: %s", domain, err)
		return
	}

	for _, ans := range r.Answer {
		record, isType := ans.(*dns.SRV)
		if isType {
			fmt.Println(record.Target, record.Port)
		}
	}
}

func QueryDNSCnameRecord(domain string) {
	c := dns.Client{Timeout: DNSTimeout * time.Second}

	m := dns.Msg{}
	m.SetQuestion(SetFQDN(domain), dns.TypeCNAME)

	r, _, err := c.Exchange(&m, DefaultDNSServer)
	if err != nil {
		log.Fatalf("query SRV record failed, domain: %s, err: %s", domain, err)
		return
	}

	for _, ans := range r.Answer {
		record, isType := ans.(*dns.CNAME)
		if isType {
			fmt.Println(record.Target)
		}
	}
}
