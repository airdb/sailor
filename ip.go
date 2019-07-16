package sailor

import (
	"log"
	"net"
	"strconv"
	"strings"
)

// Convert uint to net.IP
func Inet_ntoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// Convert net.IP to int64
func Inet_aton(ipnr string) int64 {
	// parse ip to net.IP,  return nil if not ip
	addr := net.ParseIP(ipnr)
	if addr == nil {
		log.Printf("%v is not a vaild ip address.\n", ipnr)
		return -1
	}

	bits := strings.Split(ipnr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

/*
// Convert net.IP to int
func Inet_aton(ipnr string ) int {
    // parse ip to net.IP,  return nil if not ip
    addr := net.ParseIP( ipnr )
    if addr == nil {
       log.Printf( "%v is not a vaild ip address.\n", ipnr )
       return -1
    }

    bits := strings.Split(ipnr, ".")

    b0, _ := strconv.Atoi(bits[0])
    b1, _ := strconv.Atoi(bits[1])
    b2, _ := strconv.Atoi(bits[2])
    b3, _ := strconv.Atoi(bits[3])

    var sum int

    sum += int(b0) << 24
    sum += int(b1) << 16
    sum += int(b2) << 8
    sum += int(b3)

    return sum
}
*/
