package nginx_test

import (
	"testing"

	"github.com/airdb/sailor/nginx"
)

func TestGenServerConfig(t *testing.T) {
	var config nginx.ServerConfig
	config.Listen = "80"
	// config.Server_name = "_"
	config.ServerName = "_"
	config.Log.AccessLog = "logs/$host.access.log"
	config.Log.ErrorLog = "logs/error.log"
	//config.Cert.Ssl_certificate_key = "ssl-key"
	//config.Cert.Ssl_certificate = "ssl"
	config.Cert.SSLCertificateKey = "ssl-key"
	config.Cert.SSLCertificate = "ssl"

	var location nginx.Location
	location.Key = "/"
	location.Return = "200"
	location.Index = "index.html"
	config.Locations = append(config.Locations, location)

	location.Key = "~ /hello"
	location.Return = "301"
	location.Index = "index.html"

	var proxy nginx.ProxySetHeader
	proxy.ProxySetHeader = "Host             $host"
	location.ProxySetHeaders = append(location.ProxySetHeaders, proxy)

	proxy.ProxySetHeader = "X-Real-IP        $remote_addr"
	location.ProxySetHeaders = append(location.ProxySetHeaders, proxy)

	proxy.ProxySetHeader = "X-Forwarded-For  $proxy_add_x_forwarded_for"
	location.ProxySetHeaders = append(location.ProxySetHeaders, proxy)

	location.Proxy.ProxyRedirect = "on"

	var include nginx.Include
	include.Include = "conf.d/gzip_params.conf"

	location.Includes = append(location.Includes, include)

	include.Include = "conf.d/fastcgi.conf"
	location.Includes = append(location.Includes, include)

	config.Locations = append(config.Locations, location)

	t.Log("test log xxxx", config)

	nginx.GenServerConfig(&config)
}

func TestGenUpstreamConfig(t *testing.T) {
	var config nginx.UpstreamConfig

	/*
		var ups nginx.Upstream
		ups.Server = "127.0.0.1"
		ups.BalancerByLuaBlock = "gw.set_current_peer()"
		ups.Keepalive = 32
		config.Upstreams = append(config.Upstreams, ups)
		config.Key = "hello-world-test-sg"
	*/

	var server nginx.UpstreamServer
	server.Server = "10.10.10.10:8080"
	config.Servers = append(config.Servers, server)

	server.Server = "10.11.11.11:8080 backup"
	config.Servers = append(config.Servers, server)
	config.Key = "hello-world-ip-upstream"

	t.Log("test upstream config")
	nginx.GenUpstreamConfig(&config)
}
