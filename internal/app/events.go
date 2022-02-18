package app

type Event struct {
	IsProxyEnabled bool
}

type EventListener func(e *Event)
