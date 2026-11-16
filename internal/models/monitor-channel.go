package models

type MonitorChannel struct {
	MonitorID string `gorm:"primaryKey;size:64" json:"monitor_id"`
	ChannelID string `gorm:"primaryKey;size:64" json:"channel_id"`
}

func (MonitorChannel) TableName() string { return "monitor_channels" }
