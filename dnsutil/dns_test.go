package dnsutil

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryDNSSRVRecord(t *testing.T) {
	Convey("Query DNS SRV record", t, func() {
		rr := QueryDNSSRVRecord("hello.airdb.me")

		Convey("Then trim space", func() {
			t.Log(rr.Target, rr.Port)
		})
	})
}

func TestQueryDNSCnameRecord(t *testing.T) {
	Convey("Query DNS Cname record", t, func() {
		rr := QueryDNSCnameRecord("airdb.dev")

		Convey("Then trim space", func() {
			t.Log(rr.Target)
		})
	})
}
