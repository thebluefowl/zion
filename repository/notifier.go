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

func (r *NotifierRepository) GetAll(tenantID, subscriberID string) ([]model.Notifier, error) {
	notifiers := []model.Notifier{}
	err := r.db.Preload("Subscriber").Find(&notifiers, model.Notifier{TenantID: tenantID, SubscriberID: subscriberID}).Error
	return notifiers, err
}

func (r *NotifierRepository) GetByType(tenantID, subscriberID string, notifierType model.NotifierType) ([]model.Notifier, error) {
	notifiers := []model.Notifier{}
	err := r.db.Preload("Subscriber").Find(&notifiers, model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, NotifierType: notifierType}).Error
	return notifiers, err
}

func (r *NotifierRepository) Delete(tenantID, subscriberID, id string) error {
	return r.db.Delete(&model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, ID: id}).Error
}
