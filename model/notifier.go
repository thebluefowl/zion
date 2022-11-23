package model

type NotifierType string

const NotifierTypeWebhook NotifierType = "webhook"

type NotifierClass string

const NotifierClassDefault NotifierClass = "default"

var NotifierTypeClassMap = map[NotifierType]NotifierClass{
	NotifierTypeWebhook: NotifierClassDefault,
}

type Notifier struct {
	ID           string        `gorm:"primary_key"`
	Type         NotifierType  `gorm:"column:type;not null"`
	Class        NotifierClass `gorm:"column:class;not null"`
	SubscriberID string        `gorm:"not null"`
	Subscriber   Subscriber    `gorm:"foreignkey:SubscriberID"`
	TenantID     string        `gorm:"not null"`
	Tenant       Tenant        `gorm:"foreignkey:TenantID"`
	Config       []byte        `gorm:"type:jsonb;not null"`
	IsActive     bool          `gorm:"not null;default:true"`
}

func (n *Notifier) SetClass() {
	n.Class = NotifierTypeClassMap[n.Type]
}

type NotifierRepository interface {
	Create(notifier *Notifier) error
	Get(tenantID, subscriberID, id string) (*Notifier, error)
	Filter(tenantID, subscriberID string, notifierType NotifierType, isActive bool) ([]Notifier, error)
	Delete(tenantID, subscriberID, id string) error
}
