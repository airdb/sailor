package sailor

import (
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

func QueryDNSSRVRecord(domain string) *dns.SRV {
	c := dns.Client{Timeout: DNSTimeout * time.Second}

	m := dns.Msg{}
	m.SetQuestion(SetFQDN(domain), dns.TypeSRV)

	r, _, err := c.Exchange(&m, DefaultDNSServer)
	if err != nil {
		log.Fatalf("query SRV record failed, domain: %s, err: %s", domain, err)
		return nil
	}

	for _, ans := range r.Answer {
		record, isType := ans.(*dns.SRV)
		if isType {
			return record
		}
	}

	return nil
}

func QueryDNSCnameRecord(domain string) *dns.CNAME {
	c := dns.Client{Timeout: DNSTimeout * time.Second}

	m := dns.Msg{}
	m.SetQuestion(SetFQDN(domain), dns.TypeCNAME)

	r, _, err := c.Exchange(&m, DefaultDNSServer)
	if err != nil {
		log.Fatalf("query SRV record failed, domain: %s, err: %s", domain, err)
		return nil
	}

	for _, ans := range r.Answer {
		record, isType := ans.(*dns.CNAME)
		if isType {
			return record
		}
	}

	return nil
}
