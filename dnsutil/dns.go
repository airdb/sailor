package dnsutil

import (
	"log"
	"strings"
	"time"

	"github.com/airdb/sailor"
	"github.com/miekg/dns"
	"golang.org/x/net/publicsuffix"
)

const (
	DefaultDNSServer = "8.8.8.8:53"
)

var (
	DNSTimeout = 3
	DNSRetry   = 3
)

// FQDN: A fully qualified domain name.
func SetFQDN(domain string) string {
	if !strings.HasSuffix(domain, ".") {
		domain += "."
	}

	return domain
}

func QueryDNSSRVRecord(domain string) *dns.SRV {
	c := dns.Client{Timeout: time.Duration(DNSTimeout) * time.Second}

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
	c := dns.Client{Timeout: time.Duration(DNSTimeout) * time.Second}

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

func ToDomainWithDot(domain string) string {
	if !strings.HasSuffix(domain, ".") {
		domain += "."
	}

	return domain
}

func TrimDomainDot(domain string) string {
	return strings.TrimSuffix(domain, sailor.DelimiterDot)
}

func GetRootDomain(name string) (string, error) {
	name = TrimDomainDot(name)

	rootDomain, err := publicsuffix.EffectiveTLDPlusOne(name)
	if err != nil {
		return "", err
	}

	return rootDomain, nil
}

func GetDomainSuffix(name string) string {
	name = TrimDomainDot(name)
	publicSuffix, _ := publicsuffix.PublicSuffix(name)

	return publicSuffix
}
