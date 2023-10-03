package transport

import "time"

type RetryConfig struct {
	MaxRetries      int
	InitialDelay    time.Duration
	MaxDelay        time.Duration
	ExponentialBase float64
	Jitter          bool
	RetryOnStatus   []int
}

type HTTPClientConfig struct {
	SemaphoreBufferSize int
	Timeout             time.Duration
	KeepAlive           time.Duration
	MaxIdleConns        int
	IdleConnTimeout     time.Duration
}
