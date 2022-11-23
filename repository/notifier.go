package repository

import (
	"github.com/thebluefowl/zion/model"
	"gorm.io/gorm"
)

type NotifierRepository struct {
	db *gorm.DB
}

func NewNotifierRepository(db *gorm.DB) model.NotifierRepository {
	return &NotifierRepository{db: db}
}

func (r *NotifierRepository) Create(n *model.Notifier) error {
	return r.db.Create(n).Error
}

func (r *NotifierRepository) Get(tenantID, subscriberID, id string) (*model.Notifier, error) {
	notifier := &model.Notifier{}
	err := r.db.Preload("Subscriber").First(notifier, model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, ID: id}).Error
	return notifier, err
}

func (r *NotifierRepository) Filter(tenantID, subscriberID string, notifierType model.NotifierType, isActive bool) ([]model.Notifier, error) {
	notifiers := []model.Notifier{}
	filter := model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, IsActive: isActive}
	if notifierType != "" {
		filter.NotifierType = notifierType
	}
	err := r.db.Preload("Subscriber").Preload("Tenant").Preload("Subscriber.Tenant").Find(&notifiers, filter).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return notifiers, err
}

func (r *NotifierRepository) Delete(tenantID, subscriberID, id string) error {
	return r.db.Delete(&model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, ID: id}).Error
}
