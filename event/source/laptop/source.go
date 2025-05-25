package laptop

import (
	"github.com/google/uuid"
	"goem"
	"goem/internal/event"
	"goem/internal/source"
)

const (
	emptyUUid = "00000000-0000-0000-0000-000000000000"
	unknown   = "Unknown"
)

type Laptop interface {
	GetBrand() string
	GetModel() string
	GetRam() int
	GetStorage() int
	GetCpuId() string
	GetBoardSerialNumber() string
	GetBiosSerialNumber() string
	// behaviors
	onNewed(*Newed) error
	Register() error
	onRegistered(*Registered) error
	Refresh() error
	onRefreshed(*Refreshed) error
	// EventSourcer event source
	source.EventSourcer[Laptop]
}

type laptop struct {
	brand             string
	model             string
	ram               int
	storage           int
	cpuId             string
	boardSerialNumber string
	biosSerialNumber  string
	es                source.EventHandler[Laptop]
}

func (lt *laptop) GetChanges() []event.Applier[Laptop] {
	return lt.es.GetChanges()
}

func (lt *laptop) AcceptChanges() {
	lt.es.AcceptChanges()
}

func (lt *laptop) Version() int {
	return lt.es.Version()
}

func (lt *laptop) GetBrand() string {
	return lt.brand
}

func (lt *laptop) GetModel() string {
	return lt.model
}

func (lt *laptop) GetRam() int {
	return lt.ram
}

func (lt *laptop) GetStorage() int {
	return lt.storage
}

func (lt *laptop) GetCpuId() string {
	return lt.cpuId
}

func (lt *laptop) GetBoardSerialNumber() string {
	return lt.boardSerialNumber
}

func (lt *laptop) GetBiosSerialNumber() string {
	return lt.biosSerialNumber
}

type Config struct {
	Brand   string
	Model   string
	Ram     int
	Storage int
}

func newEventSource() *laptop {
	lt := &laptop{}
	lt.es = source.NewEventSource[Laptop](lt)
	return lt
}

func NewFromEvents(events ...event.Applier[Laptop]) Laptop {
	lt := newEventSource()
	for _, e := range events {
		err := lt.es.Handle(e)
		if err != nil {
			return nil
		}
	}
	lt.AcceptChanges()
	return lt
}

func New(config Config) Laptop {
	lt := newEventSource()
	var e = &Newed{
		Brand:             config.Brand,
		Model:             config.Model,
		Ram:               config.Ram,
		Storage:           config.Storage,
		CpuId:             uuid.Nil.String(),
		BoardSerialNumber: uuid.Nil.String(),
		BiosSerialNumber:  uuid.Nil.String(),
		Event:             &event.Event[Laptop]{Version: 1, Name: "laptop.Newed"},
	}
	err := lt.es.HandleNew(e)
	if err != nil {
		return nil
	}

	return lt
}

func (lt *laptop) onNewed(e *Newed) error {
	lt.brand = e.Brand
	lt.model = e.Model
	lt.ram = e.Ram
	lt.storage = e.Storage
	lt.cpuId = e.CpuId
	lt.boardSerialNumber = e.BoardSerialNumber
	lt.biosSerialNumber = e.BiosSerialNumber
	return nil
}

func (lt *laptop) Register() error {
	if lt.cpuId == emptyUUid || lt.boardSerialNumber == emptyUUid || lt.biosSerialNumber == emptyUUid {
		e := &Registered{
			CpuId:             uuid.New().String(),
			BoardSerialNumber: uuid.New().String(),
			BiosSerialNumber:  uuid.New().String(),
			Event:             &event.Event[Laptop]{Version: 1, Name: "laptop.Registered"},
		}
		err := lt.es.HandleNew(e)
		if err != nil {
			return err
		}
		return nil
	}
	return goem.ErrEventFailure
}

func (lt *laptop) onRegistered(e *Registered) error {
	lt.cpuId = e.CpuId
	lt.boardSerialNumber = e.BoardSerialNumber
	lt.biosSerialNumber = e.BiosSerialNumber

	if lt.cpuId == emptyUUid || lt.boardSerialNumber == emptyUUid || lt.biosSerialNumber == emptyUUid {
		return goem.ErrEventFailure
	}

	return nil
}

func (lt *laptop) Refresh() error {
	e := &Registered{
		CpuId:             unknown,
		BoardSerialNumber: unknown,
		BiosSerialNumber:  unknown,
		Event:             &event.Event[Laptop]{Version: 1, Name: "laptop.Refreshed"},
	}

	if lt.cpuId == emptyUUid {
		e.CpuId = uuid.New().String()
	}
	if lt.boardSerialNumber == emptyUUid {
		e.BoardSerialNumber = uuid.New().String()
	}
	if lt.biosSerialNumber == emptyUUid {
		e.BiosSerialNumber = uuid.New().String()
	}

	err := lt.es.HandleNew(e)
	if err != nil {
		return err
	}

	return nil
}

func (lt *laptop) onRefreshed(e *Refreshed) error {
	if lt.cpuId == emptyUUid || lt.boardSerialNumber == emptyUUid || lt.biosSerialNumber == emptyUUid {
		return goem.ErrEventFailure
	}

	lt.cpuId = e.CpuId
	lt.boardSerialNumber = e.BoardSerialNumber
	lt.biosSerialNumber = e.BiosSerialNumber

	return nil
}
