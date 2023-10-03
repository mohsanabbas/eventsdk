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
