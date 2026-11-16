package models

import "time"

// Monitor represents an individual target to be monitored.
type Monitor struct {
	ID                 string     `gorm:"primaryKey;size:64" json:"id"`
	Name               string     `gorm:"size:255;not null" json:"name"`
	Type               string     `gorm:"size:32;not null" json:"type"`                    // http, keyword, port, docker
	Target             string     `gorm:"type:text;not null" json:"target"`                // URL, hostname, or container name
	Interval           int        `gorm:"default:60;not null" json:"interval"`             // in seconds, default 60
	Timeout            int        `gorm:"default:10;not null" json:"timeout"`              // in seconds, default 10
	Method             string     `gorm:"size:16;default:'GET'" json:"method,omitempty"`   // GET, POST, etc.
	Headers            string     `gorm:"type:text;default:'{}'" json:"headers,omitempty"` // JSON string map[string]string
	Body               string     `gorm:"type:text;default:''" json:"body,omitempty"`
	ExpectedStatusCode int        `gorm:"default:0" json:"expected_status_code,omitempty"` // e.g. 200, 0 means 2xx
	Keyword            string     `gorm:"type:text;default:''" json:"keyword,omitempty"`
	KeywordType        string     `gorm:"size:32;default:'contains'" json:"keyword_type,omitempty"` // contains, not_contains
	Port               int        `gorm:"default:0" json:"port,omitempty"`                          // for port monitoring
	DockerContainer    string     `gorm:"size:255;default:''" json:"docker_container,omitempty"`
	IsActive           bool       `gorm:"index;default:true;not null" json:"is_active"`
	Status             string     `gorm:"size:32;default:'pending';not null" json:"status"` // up, down, pending, paused
	LastLatencyMs      int64      `gorm:"default:0" json:"last_latency_ms"`
	LastCheckedAt      *time.Time `json:"last_checked_at,omitempty"`
	Uptime24h          float64    `gorm:"default:100.0" json:"uptime_24h"`
	Uptime30d          float64    `gorm:"default:100.0" json:"uptime_30d"`
	ChannelIDs         []string   `gorm:"-" json:"channel_ids,omitempty"` // assigned notification channels
	CreatedAt          time.Time  `gorm:"not null" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"not null" json:"updated_at"`
}

func (Monitor) TableName() string { return "monitors" }
