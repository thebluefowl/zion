package model

type Tenant struct {
	ID   string `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

type TenantRepository interface {
	Create(t *Tenant) error
	Get(id string) (*Tenant, error)
}
