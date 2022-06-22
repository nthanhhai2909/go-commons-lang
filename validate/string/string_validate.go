package string

import (
	"go-commons-lang/errors/illegal/argument"
	"go-commons-lang/stringutils"
)

func NotEmpty(str string, message string) error {
	if stringutils.IsEmpty(str) {
		return argument.Create(message)
	}
	return nil
}

func NotBlank(str string, message string) error {
	if stringutils.IsBlank(str) {
		return argument.Create(message)
	}
	return nil
}
