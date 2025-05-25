package laptop

import "goem/internal/event"

type Newed struct {
	Brand             string
	Model             string
	Ram               int
	Storage           int
	CpuId             string
	BoardSerialNumber string
	BiosSerialNumber  string
	*event.Event[Laptop]
}

func (e *Newed) Apply(lt Laptop) error {
	return lt.onNewed(e)
}
