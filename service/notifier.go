package service

import (
	"github.com/segmentio/ksuid"
	"github.com/thebluefowl/zion/model"
)

type NotifierService struct {
	notifierRepository model.NotifierRepository
	subscriberService  *SubscriberService
}

func NewNotifierService(
	notifierRepository model.NotifierRepository,
	subscriberService *SubscriberService,
) *NotifierService {
	return &NotifierService{
		notifierRepository: notifierRepository,
		subscriberService:  subscriberService,
	}
}

type AddNotifierRequest struct {
	SubscriberID string
	TenantID     string
	Type         model.NotifierType
	Provider     model.NotifierProvider
	Config       []byte
}

// AddNotifier adds a notifier to a subscriber
func (s *NotifierService) Create(request *AddNotifierRequest) (*model.Notifier, error) {
	subscriber, err := s.subscriberService.GetSubscriber(request.TenantID, request.SubscriberID)
	if err != nil {
		return nil, err
	}

	notifier := &model.Notifier{
		ID:           ksuid.New().String(),
		Subscriber:   *subscriber,
		SubscriberID: subscriber.ID,
		TenantID:     subscriber.TenantID,
		Tenant:       subscriber.Tenant,
		Type:         request.Type,
		Provider:     request.Provider,
		Config:       request.Config,
	}

	if err := s.notifierRepository.Create(notifier); err != nil {
		return nil, err
	}
	return notifier, nil
}

func (s *NotifierService) Filter(subscriberID, tenantID string, notifierType model.NotifierType, isActive bool) ([]model.Notifier, error) {
	return s.notifierRepository.Filter(tenantID, subscriberID, notifierType, isActive)
}
