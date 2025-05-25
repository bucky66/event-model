package laptop

import (
	"github.com/google/uuid"
	"goem/internal/event"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestNew(t *testing.T) {
	// arrange
	config := Config{
		Brand:   "dell",
		Model:   "xps 13",
		Ram:     64,
		Storage: 250,
	}

	// act
	l := New(config)

	// assert
	assert.NotNil(t, 1, len(l.GetChanges()))
	assert.Equal(t, l.Version(), len(l.GetChanges()))
}

func TestNewFromEvents(t *testing.T) {
	// arrange
	newed := &Newed{
		Brand:             "dell",
		Model:             "xps 13",
		Ram:               64,
		Storage:           250,
		CpuId:             emptyUUid,
		BoardSerialNumber: emptyUUid,
		BiosSerialNumber:  emptyUUid,
		Event:             &event.Event[Laptop]{Version: 1, Name: "laptop.Newed"},
	}

	registered := &Registered{
		CpuId:             uuid.NewString(),
		BoardSerialNumber: uuid.NewString(),
		BiosSerialNumber:  uuid.NewString(),
		Event:             &event.Event[Laptop]{Version: 1, Name: "laptop.Registered"},
	}

	// act
	l := NewFromEvents(newed, registered)

	// assert
	assert.NotNil(t, 0, len(l.GetChanges()))
	assert.Equal(t, 2, l.Version())
	assert.Equal(t, newed.Brand, l.GetBrand())
	assert.Equal(t, newed.Model, l.GetModel())
	assert.Equal(t, newed.Ram, l.GetRam())
	assert.Equal(t, newed.Storage, l.GetStorage())
	assert.Equal(t, registered.CpuId, l.GetCpuId())
	assert.Equal(t, registered.BoardSerialNumber, l.GetBoardSerialNumber())
	assert.Equal(t, registered.BiosSerialNumber, l.GetBiosSerialNumber())
}

func TestAcceptChanges(t *testing.T) {
	// arrange
	config := Config{
		Brand:   "dell",
		Model:   "xps 13",
		Ram:     64,
		Storage: 250,
	}
	sut := New(config)

	// act
	sut.AcceptChanges()

	// assert
	assert.Equal(t, 0, len(sut.GetChanges()))
	assert.Equal(t, "dell", sut.GetBrand())
	assert.Equal(t, "xps 13", sut.GetModel())
	assert.Equal(t, 64, sut.GetRam())
	assert.Equal(t, 250, sut.GetStorage())
	assert.Equal(t, emptyUUid, sut.GetCpuId())
	assert.Equal(t, emptyUUid, sut.GetBoardSerialNumber())
	assert.Equal(t, emptyUUid, sut.GetBiosSerialNumber())
	assert.Equal(t, 1, sut.Version())
}

func TestEvents(t *testing.T) {
	// arrange
	config := Config{
		Brand:   "dell",
		Model:   "xps 13",
		Ram:     64,
		Storage: 250,
	}
	sut := New(config)

	t.Run("Register", func(t *testing.T) {
		// act
		err := sut.Register()

		changes := sut.GetChanges()

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 2, len(sut.GetChanges()))
		assert.Equal(t, "laptop.Registered", changes[1].(*Registered).Name)

	})
}
