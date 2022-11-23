package service

import (
	"github.com/segmentio/ksuid"
	"github.com/thebluefowl/zion/model"
)

type NotifierService struct {
	notifierRepository model.NotifierRepository
	subscriberService  *SubscriberService
	tenantService      *TenantService
}

// AddNotifier adds a notifier to a subscriber
func (s *NotifierService) AddNotifier(request *AddNotifierRequest) error {
	subscriber, err := s.subscriberService.GetSubscriber(request.TenantID, request.SubscriberID)
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

	return s.notifierRepository.Create(notifier)
}

func (s *SubscriberService) GetNotifiers(subscriberID, tenantID string, notifierType model.NotifierType) ([]model.Notifier, error) {
	if notifierType == "" {
		return s.subscriberRepository.GetAllNotifiers(subscriberID, tenantID)
	}
	return s.subscriberRepository.GetNotifiersByType(subscriberID, tenantID, notifierType)
}
