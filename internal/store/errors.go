package store

import (
	"fmt"
)

type UserNotFoundError struct {
	ID string
}

func (err UserNotFoundError) Error() string {
	return fmt.Sprintf("User %v doesn't exist", err.ID)
}

type DuplicateUserError struct {
	ID string
}

func (err DuplicateUserError) Error() string {
	return fmt.Sprintf("student %v already exist", err.ID)
}
