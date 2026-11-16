package models

import "time"

type Heartbeat struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	MonitorID    string    `gorm:"size:64;index;not null" json:"monitor_id"`
	Status       string    `gorm:"size:32;not null" json:"status"` // up, down
	LatencyMs    int64     `gorm:"default:0" json:"latency_ms"`
	StatusCode   int       `gorm:"default:0" json:"status_code,omitempty"`
	ErrorMessage string    `gorm:"type:text;default:''" json:"error_message,omitempty"`
	Timestamp    time.Time `gorm:"index;not null" json:"timestamp"`
}

func (Heartbeat) TableName() string { return "heartbeats" }
