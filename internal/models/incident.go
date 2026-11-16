package models

import "time"

type Incident struct {
	ID              string     `gorm:"primaryKey;size:64" json:"id"`
	MonitorID       string     `gorm:"size:64;index;not null" json:"monitor_id"`
	MonitorName     string     `gorm:"-" json:"monitor_name,omitempty"`
	Status          string     `gorm:"size:32;not null" json:"status"` // down, resolved
	StartedAt       time.Time  `gorm:"not null" json:"started_at"`
	ResolvedAt      *time.Time `json:"resolved_at,omitempty"`
	DurationSeconds int64      `gorm:"default:0" json:"duration_seconds,omitempty"`
	Cause           string     `gorm:"type:text;default:''" json:"cause"`
}

func (Incident) TableName() string { return "incidents" }
