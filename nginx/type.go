package nginx

type Config struct {
	Daemon          string
	WorkerProcesses string
	Log
	Pid string
	Events
	HTTPConfig
}

type Events struct {
	WorkerConnections uint
}

type HTTPConfig struct {
	LogFormat string
	Log
	KeepaliveTimeout          uint
	ServerNamesHashMaxSize    uint
	ServerNamesHashBucketSize uint
	ServerTokens              string
	Sendfile                  string
	TCPNopush                 string
	TCPNodelay                string

	Client
	SendTimeout    string
	PortInRedirect string
	DefaulType     string
	Servers        []ServerConfig
	GZIP
	Proxy
	Fastcgi
	ReqStatusZone string
}

type Client struct {
	ClientMaxBodySize   string
	ClientHeaderTimeout string
	ClientBodyTimeout   string
}

type GZIP struct {
	Gzip            string
	GzipTypes       string
	GzipMinLength   string
	GzipBuffers     string
	GzipHTTPVersion string
	GzipCompLevel   string
	GzipVary        string
}

type Proxy struct {
	ProxyNextUpstream    string
	ProxySetHeaders      []ProxySetHeader
	ProxyInterceptErrors string
	ProxyRedirect        string

	ProxyFferSize         string
	ProxyBuffering        string
	ProxyBuffers          string
	ProxyBusyBuffersSize  string
	ProxyCache            string
	ProxyCacheBypass      string
	ProxyCacheKey         string
	ProxyCacheMethods     string
	ProxyCacheValid       string
	ProxyConnectTimeout   string
	ProxyCookieDomain     string
	ProxyHideHeader       string
	ProxyHTTPVersion      string
	ProxyIgnoreHeaders    string
	ProxyMethod           string
	ProxyNoCache          string
	ProxyPass             string
	ProxyReadTimeout      string
	ProxyRequestBuffering string
	ProxySendTimeout      string
	ProxySSLServerName    string
}

type ProxySetHeader struct {
	ProxySetHeader string
}

type Fastcgi struct {
	FastcgiConnectTimeout  string
	FastcgiSendTimeout     string
	FastcgiReadTimeout     string
	FastcgiBufferSize      string
	FastcgiBuffers         string
	FastcgiBusyBuffersSize string
	FastcgiInterceptErrors string
}

type Servers struct {
	// SERVERS []Server
	ServerConfig
}

type ServerConfig struct {
	Listen string
	// Server_name string
	ServerName string
	Locations  []Location
	Log
	Cert
	Includes      []Include
	ReqStatus     string
	ReqStatusShow string
}

type Include struct {
	Include string
}

type Location struct {
	Key   string
	Root  string
	Index string

	Return         string
	Includes       []Include
	Include        string
	AuthRequest    string
	AuthRequestSet string
	Rewrite        string
	ProxyPass      string
	Proxy
	Log
	TryFiles          string
	Expires           string
	Etag              string
	ClientMaxBodySize string
	Set               string
	AddHeaders        []AddHeader
	Alias             string
	AccessControl
	Lua
	ErrorPage                string
	DefaultType              string
	GrpcPass                 string
	Keepalive                string
	KeepaliveRequests        string
	KeepaliveTimeout         string
	LargeClientHeaderBuffers string

	SubFilters     []SubFilter
	SubFilterOnce  string
	SubFilterTypes string

	If string
}

type AccessControl struct {
	Allow string
	Deny  string
}

type Lua struct {
	SetByLuaBlock          string
	HeaderFilterByLuaBlock string
	BodyFilterByLuaBlock   string
	ContentByLuaBlock      string
}

type SubFilter struct {
	SubFilter string
}

type AddHeader struct {
	AddHeader string
}

type Cert struct {
	SSLCertificate    string
	SSLCertificateKey string
}

type Log struct {
	AccessLog string
	ErrorLog  string
}

type UpstreamConfig struct {
	Key       string
	Upstreams []Upstream
	Servers   []UpstreamServer
	Keepalive uint
}

type UpstreamServer struct {
	Server string
}

type Upstream struct {
	Server             string
	BalancerByLuaBlock string
	Keepalive          uint
}
