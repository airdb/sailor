package asn2ip_test

import (
	"testing"

	"github.com/airdb/sailor/asn2ip"
)

func TestRun(t *testing.T) {
	t.Log("TestRun")

	fetcher := asn2ip.NewFetcher("whois.radb.net", 43)
	ips, err := fetcher.Fetch(true, false, "15169")
	if err != nil {
		t.Error(err)
		return
	}

	for _, ipnet := range ips {
		for _, ip := range ipnet["ipv4"] {
			t.Log(ip)
		}
	}
}
