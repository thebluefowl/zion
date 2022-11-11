package channel

import (
	"github.com/thebluefowl/zion/dispatcher"
)

type ChannelType string

const ChannelTypeSlack ChannelType = "slack"
const ChannelTypeWebhook ChannelType = "webhook"

type Channel struct {
	ID   string
	Type ChannelType
}

func GetDispatcher(t ChannelType) dispatcher.Dispatcher {
	switch t {
	case ChannelTypeWebhook:
		return &dispatcher.Webhook{}
	default:
		return nil
	}
}
