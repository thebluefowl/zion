package model

type ChannelConfig struct {
}

type Subscriber struct {
	ID       string `gorm:"primary_key"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"column:email"`
	TenantID string
	Tenant   Tenant
}

type SubscriberRepository interface {
	Create(s *Subscriber) error
	Get(tenantID, id string) (*Subscriber, error)
	Delete(tenantID, id string) error
}
