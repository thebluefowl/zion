package subscription

// Connection is a struct that represents a set of channels used to send messages.
type Connection struct {
	ID         string      `gorm:"primary_key"`
	TenantID   string      `gorm:"primary_key"`
	Name       string      `gorm:"not null"`
	Channels   []Channel   `gorm:"foreignkey:ChannelID;association_foreignkey:ID"`
	Subscriber *Subscriber `gorm:"foreignkey:SubscriberID;association_foreignkey:ID"`
}

type ConnectionRepository interface {
	GetConnection(id, tenantID string) (*Connection, error)
	Create(connection *Connection) error
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
