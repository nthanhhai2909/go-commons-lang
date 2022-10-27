package string

import (
	"github.com/nthanhhai2909/go-commons-lang/errors/illegal/argument"
	"github.com/nthanhhai2909/go-commons-lang/validate"
	"testing"
)

func TestNotEmptyCase1(t *testing.T) {
	message := "my message"
	err := NotEmpty("", message)

	if err == nil {
		t.Errorf("NotEmpty %s failed when expression is false", "case 1")
	}
	if _, ok := err.(*argument.Error); !ok {
		t.Errorf("NotEmpty %s failed must return %s", "case 1", "Error")
	}

	if err.Error() != "my message" {
		t.Errorf("NotEmpty %s failed to return message: %s", "case 1", message)
	}
}

func TestNotEmptyCase2(t *testing.T) {
	err := NotEmpty("", "")

	if err == nil {
		t.Errorf("NotEmpty %s failed when expression is false", "case 2")
	}
	if _, ok := err.(*argument.Error); !ok {
		t.Errorf("NotEmpty %s failed must return %s", "case 2", "Error")
	}

	if err.Error() != validate.DefaultIsTrueExMessage {
		t.Errorf("NotEmpty %s failed to return default message %s", "case 2", validate.DefaultIsTrueExMessage)
	}
}

func TestNotEmptyCase3(t *testing.T) {
	err := NotEmpty("abc", "")
	if err != nil {
		t.Errorf("NotEmpty failed to validate expression")
	}
}

func TestNotBlankCase1(t *testing.T) {
	message := "my message"
	err := NotBlank("  ", message)

	if err == nil {
		t.Errorf("NotBlank %s failed when expression is false", "case 1")
	}
	if _, ok := err.(*argument.Error); !ok {
		t.Errorf("NotBlank %s failed must return %s", "case 1", "Error")
	}

	if err.Error() != "my message" {
		t.Errorf("NotBlank %s failed to return message: %s", "case 1", message)
	}
}

func TestNotBlankCase2(t *testing.T) {
	err := NotBlank("   ", "")

	if err == nil {
		t.Errorf("NotBlank %s failed when expression is false", "case 2")
	}
	if _, ok := err.(*argument.Error); !ok {
		t.Errorf("NotBlank %s failed must return %s", "case 2", "Error")
	}

	if err.Error() != validate.DefaultIsTrueExMessage {
		t.Errorf("NotEmpty %s failed to return default message %s", "case 2", validate.DefaultIsTrueExMessage)
	}
}

func TestNotBlankCase3(t *testing.T) {
	err := NotBlank("  abc  ", "")
	if err != nil {
		t.Errorf("NotBlank failed to validate expression")
	}
}
