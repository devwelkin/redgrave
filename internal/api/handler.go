package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/welkin/redgrave/internal/checker"
)

type PingRequest struct {
	URL     string `json:"url"`
	Timeout string `json:"timeout"`
}

type PingResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	LatencyMs  int64  `json:"latency_ms"`
	Error      string `json:"error,omitempty"`
}

func HandlePing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	timeout := 10 * time.Second
	if req.Timeout != "" {
		d, err := time.ParseDuration(req.Timeout)
		if err != nil {
			http.Error(w, "invalid timeout", http.StatusBadRequest)
			return
		}
		timeout = d
	}

	result := checker.Ping(req.URL, timeout)

	resp := PingResponse{
		Success:    result.Success,
		StatusCode: result.StatusCode,
		LatencyMs:  result.Latency.Milliseconds(),
		Error:      result.Error,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
