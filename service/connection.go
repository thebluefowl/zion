package service

import (
	"github.com/segmentio/ksuid"
	"github.com/thebluefowl/zion/subscription"
)

type ConnectionService struct {
	connectionRepository subscription.ConnectionRepository
	channelRepository    subscription.ChannelRepository
	subscriberRepository subscription.SubscriberRepository
}

func NewConnectionService(connectionRepository subscription.ConnectionRepository) *ConnectionService {
	return &ConnectionService{connectionRepository: connectionRepository}
}

type CreateConnectionRequest struct {
	TenantID     string
	Name         string
	SubscriberID string
	ChannelIDs   []string
}

func (s *ConnectionService) CreateConnection(request *CreateConnectionRequest) error {
	channels, err := s.channelRepository.FilterChannels(request.TenantID, request.ChannelIDs)
	if err != nil {
		return err
	}

	subscriber, err := s.subscriberRepository.Get(request.SubscriberID, request.TenantID)
	if err != nil {
		return err
	}

	connection := &subscription.Connection{
		ID:         ksuid.New().String(),
		TenantID:   request.TenantID,
		Name:       request.Name,
		Channels:   channels,
		Subscriber: subscriber,
	}

	return s.connectionRepository.Create(connection)
}
