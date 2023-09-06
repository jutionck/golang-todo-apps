package exception

import "errors"

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() error {
	return errors.New(v.Message)
}
