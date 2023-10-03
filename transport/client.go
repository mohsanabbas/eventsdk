package transport

import (
	"net/http"
	"time"
)

const (
	SemaphoreBufferSize = 100
	handshakeTimeout    = 10 * time.Second
)

type CustomHTTPClient struct {
	Client    *http.Client
	semaphore chan struct{}
}

func NewCustomHTTPClient(config HTTPClientConfig) *CustomHTTPClient {
	return &CustomHTTPClient{
		Client: &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				MaxIdleConns:        config.MaxIdleConns,
				IdleConnTimeout:     config.IdleConnTimeout,
				DisableKeepAlives:   false,
				TLSHandshakeTimeout: handshakeTimeout,
			},
		},
		semaphore: make(chan struct{}, config.SemaphoreBufferSize),
	}
}

func (c *CustomHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.semaphore <- struct{}{}
	defer func() { <-c.semaphore }()

	return c.Client.Do(req)
}
