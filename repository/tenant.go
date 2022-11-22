package repository

import (
	"github.com/thebluefowl/zion/model"
	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) model.TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) Create(t *model.Tenant) error {
	return r.db.Create(t).Error
}

func (r *TenantRepository) Get(id string) (*model.Tenant, error) {
	tenant := &model.Tenant{}
	err := r.db.First(tenant, model.Tenant{ID: id}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return tenant, err
}
