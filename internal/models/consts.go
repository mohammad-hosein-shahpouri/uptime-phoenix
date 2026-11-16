package models

// Monitor types
const (
	MonitorTypeHTTP    = "http"
	MonitorTypeKeyword = "keyword"
	MonitorTypePort    = "port"
	MonitorTypeDocker  = "docker"
)

// Monitor statuses
const (
	StatusUp      = "up"
	StatusDown    = "down"
	StatusPending = "pending"
	StatusPaused  = "paused"
)

// Keyword matching types
const (
	KeywordContains    = "contains"
	KeywordNotContains = "not_contains"
)

// Notification channel types
const (
	ChannelTypeTelegram = "telegram"
	ChannelTypeDiscord  = "discord"
	ChannelTypeSlack    = "slack"
	ChannelTypeEmail    = "email"
	ChannelTypeWebhook  = "webhook"
)
