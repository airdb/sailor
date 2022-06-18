package asn2ip_test

import (
	"strings"
	"testing"

	"github.com/airdb/sailor/asn2ip"
)

func TestRun(t *testing.T) {
	t.Log("TestRun")

	fetcher := asn2ip.NewFetcher("whois.radb.net", 43)

	asnum := "AS15169"
	asnum = "AS4134"
	asnum = strings.Replace(asnum, "AS", "", -1)
	ips, err := fetcher.Fetch(true, false, asnum)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("AS number %s ipv4 count %d", asnum, len(ips))
}
