package checker

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"time"
)

type Result struct {
	Success    bool
	StatusCode int
	Latency    time.Duration
	Error      string
}

// Ping sends a GET to url and returns a structured Result.
// It classifies the outcome: success (2xx), timeout, or other error.
func Ping(url string, timeout time.Duration) Result {
	client := &http.Client{
		Timeout: timeout,
	}

	start := time.Now()
	resp, err := client.Get(url)
	latency := time.Since(start)

	if err != nil {
		return Result{
			Success: false,
			Latency: latency,
			Error:   classify(err),
		}
	}
	defer resp.Body.Close()

	success := resp.StatusCode >= 200 && resp.StatusCode < 300
	return Result{
		Success:    success,
		StatusCode: resp.StatusCode,
		Latency:    latency,
	}
}

func classify(err error) string {
	if os.IsTimeout(err) {
		return "timeout"
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return "timeout"
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return "timeout"
	}
	return err.Error()
}
