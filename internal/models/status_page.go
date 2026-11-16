package models

import "time"

type StatusPage struct {
	ID                 string              `gorm:"primaryKey;size:64" json:"id"`
	Title              string              `gorm:"size:255;not null" json:"title"`
	Slug               string              `gorm:"size:255;unique;not null" json:"slug"`
	Description        string              `gorm:"type:text;default:''" json:"description"`
	CustomAnnouncement string              `gorm:"type:text;default:''" json:"custom_announcement,omitempty"`
	AnnouncementType   string              `gorm:"size:32;default:'info'" json:"announcement_type,omitempty"` // info, warning, critical, success
	IsPublished        bool                `gorm:"default:true;not null" json:"is_published"`
	Monitors           []StatusPageMonitor `gorm:"foreignKey:StatusPageID;references:ID" json:"monitors,omitempty"`
	CreatedAt          time.Time           `gorm:"not null" json:"created_at"`
	UpdatedAt          time.Time           `gorm:"not null" json:"updated_at"`
}

func (StatusPage) TableName() string { return "status_pages" }
