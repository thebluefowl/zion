package dispatcher

type Dispatcher interface {
	Send(payload []byte, responseCallback func([]byte, error) error) error
}
