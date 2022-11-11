package connection

import (
	"github.com/thebluefowl/zion/dispatcher"
)

type ChannelType string

const ChannelTypeSlack ChannelType = "slack"
const ChannelTypeWebhook ChannelType = "webhook"

type Channel struct {
	ID          string      `gorm:"primary_key"`
	TenantID    string      `gorm:"primary_key"`
	Priority    int         `gorm:"default:0"`
	Type        ChannelType `gorm:"not null"`
	Connections Connection  `gorm:"many2many:connection_channels;"`
	Config      interface{} `gorm:"type:jsonb;not null"`
}

type PublishMode string

const (
	PublishModeBroadcast PublishMode = "broadcast"
	PublishModePriority  PublishMode = "priority"
)

// Connection is a struct that represents a set of channels used to send messages.
type Connection struct {
	ID       string    `gorm:"primary_key"`
	TenantID string    `gorm:"primary_key"`
	Name     string    `gorm:"not null"`
	Channels []Channel `gorm:"many2many:connection_channels;"`
}

type ConnectionRepository interface {
	GetConnection(id, tenantID string) (*Connection, error)
	CreateConnection(s *Connection) error
	GetChannels(id, tenantID string) ([]Channel, error)
}

func AddChannel(s *Connection, c Channel) {
	s.Channels = append(s.Channels, c)
}

func RemoveChannel(s *Connection, c Channel) {
	for i, v := range s.Channels {
		if v.ID == c.ID {
			s.Channels = append(s.Channels[:i], s.Channels[i+1:]...)
		}
	}
}

func GetDispatcher(t ChannelType, config interface{}) dispatcher.Dispatcher {
	switch t {
	case ChannelTypeWebhook:
		return dispatcher.NewWebhook(config)
	default:
		return nil
	}
}
