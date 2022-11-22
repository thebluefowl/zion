package service

import (
	"github.com/segmentio/ksuid"
	"github.com/thebluefowl/zion/model"
)

type SubscriberService struct {
	subscriberRepository model.SubscriberRepository
	tenantService        *TenantService
}

func NewService(subscriberRepository model.SubscriberRepository) *SubscriberService {
	return &SubscriberService{subscriberRepository: subscriberRepository}
}

type CreateRequest struct {
	Name     string
	Email    string
	TenantID string
}

func (s *SubscriberService) CreateSubscriber(request *CreateRequest) error {
	tenant, err := s.tenantService.Get(request.TenantID)
	if err != nil {
		return err
	}
	subscriber := &model.Subscriber{
		ID:     ksuid.New().String(),
		Name:   request.Name,
		Email:  request.Email,
		Tenant: tenant,
	}
	return s.subscriberRepository.Create(subscriber)
}

type GetSubscriberRequest struct {
	ID       string
	TenantID string
}

func (s *SubscriberService) GetSubscriber(request *GetSubscriberRequest) (*model.Subscriber, error) {
	return s.subscriberRepository.Get(request.ID, request.TenantID)
}

type AddNotifierRequest struct {
	SubscriberID string
	TenantID     string
	NotifierType model.NotifierType
	Config       map[string]interface{}
}

func (s *SubscriberService) AddNotifier(request *AddNotifierRequest) error {
	subscriber, err := s.subscriberRepository.Get(request.SubscriberID, request.TenantID)
	if err != nil {
		return err
	}
	notifier := &model.Notifier{
		ID:         ksuid.New().String(),
		Type:       request.NotifierType,
		Config:     request.Config,
		Subscriber: *subscriber,
	}
	subscriber.Notifiers = append(subscriber.Notifiers, *notifier)
	return s.subscriberRepository.Update(subscriber)
}

func (s *SubscriberService) GetNotifier(subscriberID, tenantID, notifierID string) (*model.Notifier, error) {
	return s.subscriberRepository.GetNotifier(subscriberID, tenantID, notifierID)
}

func (s *SubscriberService) GetNotifiersByType(subscriberID, tenantID string, notifierType model.NotifierType) ([]model.Notifier, error) {
	return s.subscriberRepository.GetNotifiersByType(subscriberID, tenantID, notifierType)
}

func (s *SubscriberService) DeleteNotifier(subscriberID, tenantID string, notifierID string) error {
	return s.subscriberRepository.DeleteNotifier(subscriberID, tenantID, notifierID)
}
