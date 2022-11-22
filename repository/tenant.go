package repository

import (
	"github.com/thebluefowl/zion/model"
	"gorm.io/gorm"
)

type TenantPGRepository struct {
	db *gorm.DB
}

func NewTenantPGRepository(db *gorm.DB) *TenantPGRepository {
	return &TenantPGRepository{db: db}
}

func (r *TenantPGRepository) Create(t *model.Tenant) error {
	return r.db.Create(t).Error
}

func (r *TenantPGRepository) Get(id string) (*model.Tenant, error) {
	tenant := &model.Tenant{}
	err := r.db.First(tenant, model.Tenant{ID: id}).Error
	return tenant, err
}
