package domain

import "time"

type MonitorStatus string

const (
	MonitorStatusUp     MonitorStatus = "up"
	MonitorStatusDown   MonitorStatus = "down"
	MonitorStatusPaused MonitorStatus = "paused"
)

type Monitor struct {
	ID        string
	Name      string
	URL       string
	Interval  time.Duration
	Timeout   time.Duration
	Status    MonitorStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
