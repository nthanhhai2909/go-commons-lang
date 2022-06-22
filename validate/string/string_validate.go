package string

import (
	"go-commons-lang/errors/illegal"
	"go-commons-lang/stringutils"
)

func NotEmpty(str string, message string) error {
	if stringutils.IsEmpty(str) {
		return illegal.CreateIllegalArgumentError(message)
	}
	return nil
}

func NotBlank(str string, message string) error {
	if stringutils.IsBlank(str) {
		return illegal.CreateIllegalArgumentError(message)
	}
	return nil
}
