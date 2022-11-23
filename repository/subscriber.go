package repository

import (
	"github.com/thebluefowl/zion/model"
	"gorm.io/gorm"
)

type SubscriberRepository struct {
	db *gorm.DB
}

func NewSubscriberRepository(db *gorm.DB) model.SubscriberRepository {
	return &SubscriberRepository{db: db}
}

func (r *SubscriberRepository) Create(s *model.Subscriber) error {
	return r.db.Create(s).Error
}

func (r *SubscriberRepository) Get(tenantID, id string) (*model.Subscriber, error) {
	subscriber := &model.Subscriber{}
	err := r.db.Preload("Tenant").First(subscriber, model.Subscriber{TenantID: tenantID, ID: id}).Error
	return subscriber, err
}

func (r *SubscriberRepository) Delete(tenantID, id string) error {
	return r.db.Delete(&model.Subscriber{TenantID: tenantID, ID: id}).Error
}
