package boolean

import "go-commons-lang/errors/illegal"

func IsTrue(expression bool, message string) error {
	if !expression {
		return illegal.CreateIllegalArgumentError(message)
	}
	return nil
}
