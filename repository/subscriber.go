package repository

import (
	"github.com/thebluefowl/zion/model"
	"gorm.io/gorm"
)

type SubscriberPGRepository struct {
	db *gorm.DB
}

func NewSubscriberPGRepository(db *gorm.DB) model.SubscriberRepository {
	return &SubscriberPGRepository{db: db}
}

func (r *SubscriberPGRepository) Create(s *model.Subscriber) error {
	return r.db.Create(s).Error
}

func (r *SubscriberPGRepository) Get(tenantID, id string) (*model.Subscriber, error) {
	subscriber := &model.Subscriber{}
	err := r.db.First(subscriber, model.Subscriber{TenantID: tenantID, ID: id}).Error
	return subscriber, err
}

func (r *SubscriberPGRepository) Delete(tenantID, id string) error {
	return r.db.Delete(&model.Subscriber{TenantID: tenantID, ID: id}).Error
}

func (r *SubscriberPGRepository) CreateNotifier(n *model.Notifier) error {
	return r.db.Create(n).Error
}

func (r *SubscriberPGRepository) GetNotifier(tenantID, subscriberID, id string) (*model.Notifier, error) {
	notifier := &model.Notifier{}
	err := r.db.Preload("Subscriber").First(notifier, model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, ID: id}).Error
	return notifier, err
}

func (r *SubscriberPGRepository) DeleteNotifier(tenantID, subscriberID, id string) error {
	return r.db.Delete(&model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, ID: id}).Error
}

func (r *SubscriberPGRepository) GetNotifierByType(tenantID, subscriberID string, notifierType model.NotifierType) (*model.Notifier, error) {
	notifier := &model.Notifier{}
	err := r.db.First(notifier, model.Notifier{TenantID: tenantID, SubscriberID: subscriberID, NotifierType: notifierType}).Error
	return notifier, err
}
