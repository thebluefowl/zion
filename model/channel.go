package model

type Channel struct {
	ID              string        `gorm:"primary_key"`
	Name            string        `gorm:"not null"`
	Notifiers       []Notifier    `gorm:"many2many:channel_notifiers;"`
	RoutingPolicyID string        `gorm:"not null"`
	RoutingPolicy   RoutingPolicy `gorm:"foreignkey:RoutingPolicyID"`
	TenantID        string        `gorm:"not null"`
	Tenant          Tenant        `gorm:"foreignkey:TenantID"`
}

type RoutingPolicy struct {
	ID string `gorm:"primary_key"`
}

type ChannelRepository interface {
	Create(channel *Channel) error
	Get(tenantID, id string) (*Channel, error)
	Filter(tenantID string, Name string) ([]Channel, error)
	Delete(tenantID, id string) error
}
