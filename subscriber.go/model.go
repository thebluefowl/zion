package subcriber

import (
	"github.com/thebluefowl/zion/channel"
)

type Subscriber struct {
	ID       string
	Name     string
	Channels []channel.Channel
}

func AddChannel(s *Subscriber, c channel.Channel) {
	s.Channels = append(s.Channels, c)
}

func RemoveChannel(s *Subscriber, c channel.Channel) {
	for i, v := range s.Channels {
		if v.ID == c.ID {
			s.Channels = append(s.Channels[:i], s.Channels[i+1:]...)
		}
	}
}
