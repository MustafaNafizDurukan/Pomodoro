package home

type Event struct {
	IsProxyEnabled bool
}

type EventListener func(e *Event)
