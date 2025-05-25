package laptop

import "goem/internal/event"

type Registered struct {
	CpuId             string
	BoardSerialNumber string
	BiosSerialNumber  string
	*event.Event[Laptop]
}

func (e *Registered) Apply(l Laptop) error {
	return l.onRegistered(e)
}
