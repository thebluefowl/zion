package event

type Event struct {
	Payload      map[string]interface{}
	ConnectionID string
}
