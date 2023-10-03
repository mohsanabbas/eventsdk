package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type HTTPTransport struct {
	Client    *CustomHTTPClient
	RetryConf RetryConfig
}

func (h *HTTPTransport) Send(data []byte, endpoint string) error {
	delay := h.RetryConf.InitialDelay
	var resp *http.Response
	var err error

	for i := 0; i <= h.RetryConf.MaxRetries; i++ {
		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
		if err != nil {
			return err
		}

		resp, err = h.Client.Do(req)
		if err == nil && !shouldRetry(resp.StatusCode, h.RetryConf.RetryOnStatus) {
			break
		}

		if i == h.RetryConf.MaxRetries {
			return err
		}

		if h.RetryConf.Jitter {
			delayWithJitter := delay + time.Duration(rand.Float64()*float64(delay))
			<-time.After(delayWithJitter)
		} else {
			<-time.After(delay)
		}

		delay = time.Duration(float64(delay) * h.RetryConf.ExponentialBase)
		if delay > h.RetryConf.MaxDelay {
			delay = h.RetryConf.MaxDelay
		}
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling response to JSON: %v", err)
	}

	fmt.Println(string(jsonData))

	return nil
}

func shouldRetry(statusCode int, retryStatus []int) bool {
	for _, retrySC := range retryStatus {
		if statusCode == retrySC {
			return true
		}
	}
	return false
}
