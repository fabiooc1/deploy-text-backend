package input

import (
	"errors"
)

type CreateUserInput struct {
	Name     string
	Username string
	Password string
}

func (c *CreateUserInput) Validate() error {
	if c == nil {
		return errors.New("body cannot be empty")
	}

	return nil
}
