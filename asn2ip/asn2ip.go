package asn2ip

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
)

const (
	IPVersion4 = 4
	IPVersion6 = 6
)

// https://github.com/g0dsCookie/asn2ip
// https://www.radb.net/query/help
type Fetcher interface {
	Fetch(ipv4, ipv6 bool, asn ...string) (map[string]map[string][]*net.IPNet, error)
}

type fetcher struct {
	host string
	port int
}

type cachedFetcher struct {
	*fetcher
}

func NewFetcher(host string, port int) Fetcher {
	return &fetcher{
		host: host,
		port: port,
	}
}

func NewCachedFetcher(host string, port int) Fetcher {
	return &cachedFetcher{
		fetcher: &fetcher{host: host, port: port},
	}
}

func readLine(conn net.Conn) (string, error) {
	resp := bytes.Buffer{}
	buf := make([]byte, 1)
	for {
		if _, err := conn.Read(buf[:1]); err != nil {
			return "", errors.Wrap(err, "failed to read next byte from connection")
		}
		if buf[0] == '\n' {
			break
		} else {
			if _, err := resp.Write(buf); err != nil {
				return "", errors.Wrap(err, "failed to write received byte to buffer")
			}
		}
	}
	return strings.TrimRight(resp.String(), "\r"), nil
}

// more command refer here, https://www.radb.net/query/help
func fetch(conn net.Conn, as string, version int) ([]*net.IPNet, error) {
	cmd := ""

	switch version {
	case IPVersion4:
		cmd = fmt.Sprintf("!gAS%s\n", as)
	case IPVersion6:
		cmd = fmt.Sprintf("!6AS%s\n", as)
	default:
		return nil, errors.Errorf("unknown ip protocol version %d", version)
	}

	if _, err := conn.Write([]byte(cmd)); err != nil {
		return nil, errors.Wrapf(err, "failed to fetch ip addresses for %s", as)
	}

	response := []*net.IPNet{}
	state := "start"
	for {
		line, err := readLine(conn)
		if err != nil {
			panic(err) // TODO
		}

		if line == "D" {
			return nil, errors.Errorf("as %s not found", as)
		} else if line == "C" {
			return response, nil
		}

		if state == "start" {
			if len(line) == 0 {
				return nil, errors.Errorf("empty response for as %s", as)
			}
			if line[0] != 'A' {
				return nil, errors.Errorf("received invalid response for as %s", as)
			}
			state = "response"
			continue
		} else if state == "response" {
			nets := strings.Split(line, " ")
			for _, n := range nets {
				_, net, err := net.ParseCIDR(n)
				if err != nil {
					return nil, errors.Errorf("failed to parse network %s for as %s", n, as)
				}
				response = append(response, net)
			}
		}
	}
}

func (f *fetcher) Fetch(ipv4, ipv6 bool, asn ...string) (map[string]map[string][]*net.IPNet, error) {
	result := map[string]map[string][]*net.IPNet{}
	if len(asn) == 0 {
		return result, nil
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", f.host, f.port))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect to %s:%d", f.host, f.port)
	}
	defer func() {
		// gracefully close socket
		_, err := conn.Write([]byte("exit\n"))
		if err != nil {
			panic(err) // TODO
		}

		conn.Close()
	}()

	// enable multiple commands per connection
	if _, err := conn.Write([]byte("!!\n")); err != nil {
		return nil, errors.Wrapf(err, "failed to enable multicommand mode")
	}

	for _, v := range asn {
		result[v] = map[string][]*net.IPNet{"ipv4": {}, "ipv6": {}}
		if ipv4 {
			net, err := fetch(conn, v, IPVersion4)
			if err != nil {
				return nil, err
			}
			result[v]["ipv4"] = net
		}
		if ipv6 {
			net, err := fetch(conn, v, IPVersion6)
			if err != nil {
				return nil, err
			}
			result[v]["ipv6"] = net
		}
	}

	return result, nil
}

func (f *cachedFetcher) Fetch(ipv4, ipv6 bool, asn ...string) (map[string]map[string][]*net.IPNet, error) {
	result := map[string]map[string][]*net.IPNet{}
	if len(asn) == 0 {
		return result, nil
	}

	fmt.Println("xxx", result)
	uncached := []string{}
	for _, as := range asn {
		// fmt.Println(as)
		result[as] = map[string][]*net.IPNet{"ipv4": {}, "ipv6": {}}
		/*
			if err != nil {
				return nil, errors.Wrapf(err, "failed to fetch asn %s from cache", as)
			}

			result[as] = map[string][]*net.IPNet{"ipv4": {}, "ipv6": {}}
			if ipv4 {
				cpy := make([]*net.IPNet, 0, len(r.IPv4))
				copy(cpy, r.IPv4)
				result[as]["ipv4"] = r.IPv4
			}
			if ipv6 {
				cpy := make([]*net.IPNet, 0, len(r.IPv6))
				copy(cpy, r.IPv6)
				result[as]["ipv6"] = r.IPv6
			}
		*/
	}

	/*
		if len(uncached) == 0 {
			// all ASNs were cached
			return result, nil
		}
	*/

	// request the rest
	r, err := f.fetcher.Fetch(ipv4, ipv6, uncached...)
	if err != nil {
		return nil, err
	}

	// now cache them and append them the results
	for as, v := range r {
		if err != nil {
			return nil, errors.Wrapf(err, "failed to put %s on cache", as)
		}
		result[as] = map[string][]*net.IPNet{"ipv4": v["ipv4"], "ipv6": v["ipv6"]}
	}

	return result, nil
}
