package event

type Event interface {
	GetEntityID() string
	GetStringMessage() (string, error)
}
