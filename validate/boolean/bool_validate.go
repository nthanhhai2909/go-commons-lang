package boolean

import (
	"go-commons-lang/errors/illegal/argument"
)

func IsTrue(expression bool, message string) error {
	if !expression {
		return argument.Create(message)
	}
	return nil
}
