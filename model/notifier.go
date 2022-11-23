package model

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
type NotifierRepository interface {
	Create(notifier *Notifier) error
	Get(tenantID, subscriberID, id string) (*Notifier, error)
	Filter(tenantID, subscriberID string, notifierType NotifierType, isActive bool) ([]Notifier, error)
	Delete(tenantID, subscriberID, id string) error
}
