package illegal

import (
	"go-commons-lang/errors"
	"go-commons-lang/stringutils"
)

/**
 * The error to indicate that a method has been passed an illegal or inappropriate argument.
 */

type ArgumentError struct {
	Message string
}

func (e *ArgumentError) Error() string {
	return e.Message
}

func NewIllegalArgumentError(message string) error {
	return &ArgumentError{message}
}

func DefIllegalArgumentError() error {
	return &ArgumentError{errors.DefaultIsTrueExMessage}
}

func CreateIllegalArgumentError(message string) error {
	if stringutils.IsNotBlank(message) {
		return NewIllegalArgumentError(message)
	} else {
		return DefIllegalArgumentError()
	}
}
