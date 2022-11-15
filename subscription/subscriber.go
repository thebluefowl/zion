package subscription

type Subscriber struct {
	ID       string    `gorm:"primary_key"`
	TenantID string    `gorm:"primary_key"`
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"column:email"`
	Channels []Channel `gorm:"many2many:subscriber_channels;"`
}

func (s *Subscriber) AddChannel(c Channel) {
	s.Channels = append(s.Channels, c)
}

func (s *Subscriber) RemoveChannel(c Channel) {
	for i, v := range s.Channels {
		if v.ID == c.ID {
			s.Channels = append(s.Channels[:i], s.Channels[i+1:]...)
		}
	}
}

type SubscriberRepository interface {
	Create(s *Subscriber) error
	Update(s *Subscriber) error
	Get(id, tenantID string) (*Subscriber, error)
}
