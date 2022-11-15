package subscription

import (
	"github.com/thebluefowl/zion/dispatcher"
)

type ChannelType string

const ChannelTypeSlack ChannelType = "slack"
const ChannelTypeWebhook ChannelType = "webhook"

type Channel struct {
	ID       string      `gorm:"primary_key"`
	TenantID string      `gorm:"primary_key"`
	Priority int         `gorm:"default:0"`
	Type     ChannelType `gorm:"not null"`
	Config   interface{} `gorm:"type:jsonb;not null"`
}

type PublishMode string

const (
	PublishModeBroadcast PublishMode = "broadcast"
	PublishModePriority  PublishMode = "priority"
)

type ChannelRepository interface {
	Create(c *Channel) error
	Get(id, tenantID string) (*Channel, error)
	GetChannels(tenantID string) ([]Channel, error)
	FilterChannels(tenantID string, ids []string) ([]Channel, error)
}

func GetDispatcher(t ChannelType, config interface{}) dispatcher.Dispatcher {
	switch t {
	case ChannelTypeWebhook:
		return dispatcher.NewWebhook(config)
	default:
		return nil
	}
}
