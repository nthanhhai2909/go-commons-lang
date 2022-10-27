package string

import (
	"github.com/nthanhhai2909/go-commons-lang/errors/illegal/argument"
	"github.com/nthanhhai2909/go-commons-lang/stringutils"
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
