package service

import (
	"errors"

	"github.com/segmentio/ksuid"
	"github.com/thebluefowl/zion/subscription"
)

type SubscriberService struct {
	subscriberRepository subscription.SubscriberRepository
}

func NewService(subscriberRepository subscription.SubscriberRepository) *SubscriberService {
	return &SubscriberService{subscriberRepository: subscriberRepository}
}

type CreateSubscriberRequest struct {
	TenantID string
	Name     string
	Email    string
}

func (s *SubscriberService) CreateSubscriber(request *CreateSubscriberRequest) error {
	subscriber := &subscription.Subscriber{
		TenantID: request.TenantID,
		Name:     request.Name,
		Email:    request.Email,
	}

	return s.subscriberRepository.Create(subscriber)
}

type CreateChannelRequest struct {
	Priority int
	Type     subscription.ChannelType
	Config   interface{}
}

func (s *SubscriberService) AddChannel(id, tenantID string, c *CreateChannelRequest) (*subscription.Subscriber, error) {
	if id == "" || tenantID == "" {
		return nil, errors.New("invalid id or tenantID")
	}
	if c == nil {
		return nil, errors.New("invalid channel")
	}
	subscriber, err := s.subscriberRepository.Get(id, tenantID)
	if err != nil {
		return nil, err
	}
	channel := &subscription.Channel{
		ID:       ksuid.New().String(),
		TenantID: tenantID,
		Priority: c.Priority,
		Type:     c.Type,
		Config:   c.Config,
	}
	subscriber.AddChannel(*channel)
	return subscriber, s.subscriberRepository.Update(subscriber)
}

func (s *SubscriberService) RemoveChannel(id, tenantID, channelID string) (*subscription.Subscriber, error) {
	if id == "" || tenantID == "" || channelID == "" {
		return nil, errors.New("invalid id or tenantID or channelID")
	}
	subscriber, err := s.subscriberRepository.Get(id, tenantID)
	if err != nil {
		return nil, err
	}
	channel := subscription.Channel{
		ID:       channelID,
		TenantID: tenantID,
	}
	subscriber.RemoveChannel(channel)
	return subscriber, s.subscriberRepository.Update(subscriber)
}
