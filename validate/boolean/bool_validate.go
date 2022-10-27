package boolean

import (
	"github.com/nthanhhai2909/go-commons-lang/errors/illegal/argument"
)

func IsTrue(expression bool, message string) error {
	if !expression {
		return argument.Create(message)
	}
	return nil
}

func IsFalse(expression bool, message string) error {
	if expression {
		return argument.Create(message)
	}
	return nil
}
