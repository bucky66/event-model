package goem

type DomainError string

func (e DomainError) Error() string {
	return string(e)
}

const (
	ErrEventFailure = DomainError("event failure")
)
