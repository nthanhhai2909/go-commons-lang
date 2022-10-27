package argument

import (
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
	"github.com/nthanhhai2909/go-commons-lang/validate"
)

/**
 * The error to indicate that a method has been passed an illegal or inappropriate argument.
 */

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func New(message string) error {
	return &Error{message}
}

func Default() error {
	return &Error{validate.DefaultIsTrueExMessage}
}

func Create(message string) error {
	if stringutils.IsNotBlank(message) {
		return New(message)
	} else {
		return Default()
	}
}
