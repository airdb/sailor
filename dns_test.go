package sailor

import (
	"testing"
)

func TestQueryDNSSRVRecord(t *testing.T) {
	QueryDNSSRVRecord("hello.airdb.me")
}

func TestQueryDNSCnameRecord(t *testing.T) {
	QueryDNSCnameRecord("airdb.dev")
}
