package domain

import "github.com/google/uuid"

type SomeService interface {
	DoSomeAction(sd SomeData) error
}

type SomeData struct {
	ID uuid.UUID
}
