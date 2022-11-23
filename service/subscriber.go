package service

import (
	"github.com/segmentio/ksuid"
	"github.com/thebluefowl/zion/model"
)

type SubscriberService struct {
	subscriberRepository model.SubscriberRepository
	tenantService        *TenantService
}

func NewSubscriberService(
	subscriberRepository model.SubscriberRepository,
	tenantService *TenantService,
) *SubscriberService {
	return &SubscriberService{
		subscriberRepository: subscriberRepository,
		tenantService:        tenantService,
	}
}

type CreateRequest struct {
	Name     string
	Email    string
	TenantID string
}

func (s *SubscriberService) Create(request *CreateRequest) (*model.Subscriber, error) {
	tenant, err := s.tenantService.Get(request.TenantID)
	if err != nil {
		return nil, err
	}
	subscriber := &model.Subscriber{
		ID:       ksuid.New().String(),
		Name:     request.Name,
		Email:    request.Email,
		Tenant:   *tenant,
		TenantID: request.TenantID,
	}
	if err := s.subscriberRepository.Create(subscriber); err != nil {
		return nil, err
	}
	return subscriber, nil
}

// GetNotifiers returns all notifiers for a subscriber
func (s *SubscriberService) GetSubscriber(tenantID, id string) (*model.Subscriber, error) {
	return s.subscriberRepository.Get(tenantID, id)
}

type AddNotifierRequest struct {
	SubscriberID string
	TenantID     string
	NotifierType model.NotifierType
	Config       []byte
}

// AddNotifier adds a notifier to a subscriber
func (s *SubscriberService) AddNotifier(request *AddNotifierRequest) error {
	subscriber, err := s.subscriberRepository.Get(request.TenantID, request.SubscriberID)
	if err != nil {
		return err
	}
	notifier := &model.Notifier{
		ID:           ksuid.New().String(),
		Subscriber:   *subscriber,
		SubscriberID: subscriber.ID,
		TenantID:     subscriber.TenantID,
		Tenant:       subscriber.Tenant,
		NotifierType: request.NotifierType,
		Config:       request.Config,
	}

	return s.subscriberRepository.CreateNotifier(notifier)
}

func (s *SubscriberService) GetNotifiers(subscriberID, tenantID string, notifierType model.NotifierType) ([]model.Notifier, error) {
	if notifierType == "" {
		return s.subscriberRepository.GetAllNotifiers(subscriberID, tenantID)
	}
	return s.subscriberRepository.GetNotifiersByType(subscriberID, tenantID, notifierType)
}
