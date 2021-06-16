package acs

import (
	"log"
	"net/http"
	"time"
)

type Client struct {
	// isInsecure bool
	// regionID   string
	// config     *Config
	// httpProxy  string
	// httpsProxy string
	// noProxy    string
	// logger         *Logger
	// userAgent map[string]string
	// signer         auth.Signer
	// httpClient *http.Client
	// asyncTaskQueue chan func()
	// readTimeout    time.Duration
	// connectTimeout time.Duration
	// EndpointMap    map[string]string
	// EndpointType   string
	// Network        string
	Domain string
	// isOpenAsync    bool
}

type Config struct {
	EnableAsync       bool   `default:"false"`
	AutoRetry         bool   `default:"false"`
	Debug             bool   `default:"false"`
	MaxRetryTime      int    `default:"3"`
	MaxTaskQueueSize  int    `default:"1000"`
	GoRoutinePoolSize int    `default:"5"`
	UserAgent         string `default:""`
	Scheme            string `default:"HTTP"`
	Timeout           time.Duration
	Transport         http.RoundTripper `default:""`
	HTTPTransport     *http.Transport   `default:""`
}

func NewClientWithoutConfig() (client *Client, err error) {
	client = &Client{}

	return
}

func NewClient() (client *Client, err error) {
	client = &Client{}

	return
}

func NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret string) (client *Client, err error) {
	client = &Client{}

	log.Println(regionID, accessKeyID, accessKeySecret)

	return
}
