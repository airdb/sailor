package dnsutil

import (
	"log"
	"testing"
)

func TestQueryDNSSRVRecord(t *testing.T) {
	rr := QueryDNSSRVRecord("hello.airdb.me")
	log.Println(rr.Target, rr.Port)
}

func TestQueryDNSCnameRecord(t *testing.T) {
	rr := QueryDNSCnameRecord("airdb.dev")
	log.Println(rr.Target)
}
