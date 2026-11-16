package models

import "time"

type NotificationChannel struct {
	ID        string    `gorm:"primaryKey;size:64" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Type      string    `gorm:"size:32;not null" json:"type"`             // telegram, discord, slack, email, webhook
	Config    string    `gorm:"type:text;not null" json:"config"`         // JSON string of channel-specific config
	IsDefault bool      `gorm:"default:false;not null" json:"is_default"` // Send notifications to this channel for all monitors
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

func (NotificationChannel) TableName() string { return "notification_channels" }
