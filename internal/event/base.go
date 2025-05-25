package event

import (
	"fmt"
)

type Applier[T any] interface {
	Apply(T) error
}

type Event[TSource any] struct {
	Name    string
	Version int
}

func (e Event[TSource]) Apply(_ TSource) error {
	fmt.Println(fmt.Sprintf("apply %s event", e.Name))
	// This method is intentionally left empty.
	// It can be overridden by specific event types to apply changes to the EventSource.
	// The base implementation does nothing, allowing derived events to implement their own logic.
	return nil
}
