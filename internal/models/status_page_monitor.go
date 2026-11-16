package models

type StatusPageMonitor struct {
	StatusPageID string `gorm:"primaryKey;size:64" json:"status_page_id"`
	MonitorID    string `gorm:"primaryKey;size:64" json:"monitor_id"`
	DisplayName  string `gorm:"size:255;default:''" json:"display_name,omitempty"`
	GroupName    string `gorm:"size:255;default:''" json:"group_name,omitempty"` // e.g., "Websites", "Databases", "APIs"
	SortOrder    int    `gorm:"default:0" json:"sort_order"`
	// Populated for rendering
	Monitor *Monitor `gorm:"foreignKey:MonitorID;references:ID" json:"monitor,omitempty"`
}

func (StatusPageMonitor) TableName() string { return "status_page_monitors" }
