package model

type NotifierType string

const NotifierTypeWebhook NotifierType = "webhook"

type NotifierProvider string

const NotifierProviderDefault NotifierProvider = "default"

func ValidateNotifierType(notifierType NotifierType) bool {
	validNotifierTypes := []NotifierType{NotifierTypeWebhook}
	for _, validNotifierType := range validNotifierTypes {
		if validNotifierType == notifierType {
			return true
		}
	}
	return false
}

func ValidateNotifierProvider(notifierProvider NotifierProvider) bool {
	validNotifierProviders := []NotifierProvider{NotifierProviderDefault}
	for _, validNotifierProvider := range validNotifierProviders {
		if validNotifierProvider == notifierProvider {
			return true
		}
	}
	return false
}

type Notifier struct {
	ID           string           `gorm:"primary_key"`
	Type         NotifierType     `gorm:"column:type;not null"`
	Provider     NotifierProvider `gorm:"column:provider;not null"`
	SubscriberID string           `gorm:"not null"`
	Subscriber   Subscriber       `gorm:"foreignkey:SubscriberID"`
	TenantID     string           `gorm:"not null"`
	Tenant       Tenant           `gorm:"foreignkey:TenantID"`
	Config       []byte           `gorm:"type:jsonb;not null"`
	IsActive     bool             `gorm:"not null;default:true"`
}
type NotifierRepository interface {
	Create(notifier *Notifier) error
	Get(tenantID, subscriberID, id string) (*Notifier, error)
	Filter(tenantID, subscriberID string, notifierType NotifierType, isActive bool) ([]Notifier, error)
	Delete(tenantID, subscriberID, id string) error
}
