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
	NotifierType model.NotifierType
	Config       []byte
}

// AddNotifier adds a notifier to a subscriber
func (s *NotifierService) Create(request *AddNotifierRequest) error {
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

func (s *NotifierService) Filter(subscriberID, tenantID string, notifierType model.NotifierType, isActive bool) ([]model.Notifier, error) {
	return s.notifierRepository.Filter(tenantID, subscriberID, notifierType, isActive)
}
