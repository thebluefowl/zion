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

func (r *SubscriberRepository) CreateNotifier(n *model.Notifier) error {
	return r.db.Create(n).Error
}

func (r *SubscriberRepository) GetNotifier(tenantID, subscriberID, id string) (*model.Notifier, error) {
	notifier := &model.Notifier{}
	err := r.db.Preload("Subscriber").First(notifier, model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, ID: id}).Error
	return notifier, err
}

func (r *SubscriberRepository) GetAllNotifiers(tenantID, subscriberID string) ([]model.Notifier, error) {
	notifiers := []model.Notifier{}
	err := r.db.Preload("Subscriber").Find(&notifiers, model.Notifier{TenantID: tenantID, SubscriberID: subscriberID}).Error
	return notifiers, err
}

func (r *SubscriberRepository) DeleteNotifier(tenantID, subscriberID, id string) error {
	return r.db.Delete(&model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, ID: id}).Error
}

func (r *SubscriberRepository) GetNotifiersByType(tenantID, subscriberID string, notifierType model.NotifierType) ([]model.Notifier, error) {
	notifiers := []model.Notifier{}
	err := r.db.Preload("Subscriber").Find(&notifiers, model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, NotifierType: notifierType}).Error
	return notifiers, err
}
