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

type NotifierType string

type Notifier struct {
	ID           string       `gorm:"primary_key"`
	NotifierType NotifierType `gorm:"not null"`
	SubscriberID string       `gorm:"not null"`
	Subscriber   Subscriber   `gorm:"foreignkey:SubscriberID"`
	TenantID     string       `gorm:"not null"`
	Tenant       Tenant       `gorm:"foreignkey:TenantID"`
	Config       []byte       `gorm:"type:jsonb;not null"`
	IsActive     bool         `gorm:"not null;default:true"`
}

type SubscriberRepository interface {
	Create(s *Subscriber) error
	Get(tenantID, id string) (*Subscriber, error)
	Delete(tenantID, id string) error

	CreateNotifier(notifier *Notifier) error
	GetNotifier(tenantID, subscriberID, id string) (*Notifier, error)
	GetAllNotifiers(tenantID, subscriberID string) ([]Notifier, error)
	GetNotifiersByType(tenantID, subscriberID string, notifierType NotifierType) ([]Notifier, error)
	DeleteNotifier(tenantID, subscriberID, id string) error
}
