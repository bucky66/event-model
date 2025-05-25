package source

import (
	"goem/internal/event"
)

type EventHandler[TEventSource any] interface {
	Handle(event.Applier[TEventSource]) error
	HandleNew(event.Applier[TEventSource]) error
	EventSourcer[TEventSource]
}

type EventSourcer[TEventSource any] interface {
	GetChanges() []event.Applier[TEventSource]
	AcceptChanges()
	Version() int
}

type EventSource[TEventSource any] struct {
	changes []event.Applier[TEventSource]
	version int
	source  TEventSource
}

func NewEventSource[TEventSource any](source TEventSource) EventHandler[TEventSource] {
	return &EventSource[TEventSource]{
		source:  source,
		changes: make([]event.Applier[TEventSource], 0),
		version: 0,
	}
}

func (es *EventSource[TEventSource]) AcceptChanges() {
	es.changes = make([]event.Applier[TEventSource], 0)
}

func (es *EventSource[TEventSource]) GetChanges() []event.Applier[TEventSource] {
	return es.getChanges()
}

func (es *EventSource[TEventSource]) getChanges() []event.Applier[TEventSource] {
	if es.changes == nil {
		es.changes = []event.Applier[TEventSource]{}
	}
	return es.changes
}

func (es *EventSource[TEventSource]) Version() int {
	return es.version
}

func (es *EventSource[TEventSource]) HandleNew(e event.Applier[TEventSource]) error {
	err := es.Handle(e)
	if err != nil {
		return err
	}

	es.changes = append(es.getChanges(), e)
	return nil
}

func (es *EventSource[TEventSource]) Handle(e event.Applier[TEventSource]) error {
	err := e.Apply(es.source)
	if err != nil {
		return err
	}

	es.version += 1
	return nil
}
