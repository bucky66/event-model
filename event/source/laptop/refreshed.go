package laptop

import "goem/internal/event"

type Refreshed struct {
	CpuId             string
	BoardSerialNumber string
	BiosSerialNumber  string
	*event.Event[Laptop]
}

func (e *Refreshed) Apply(lt Laptop) error {
	return lt.onRefreshed(e)
}
