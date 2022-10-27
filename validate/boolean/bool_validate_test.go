package boolean

import (
	"github.com/nthanhhai2909/go-commons-lang/errors/illegal/argument"
	"github.com/nthanhhai2909/go-commons-lang/validate"
	"testing"
)

func TestIsTrueCase1(t *testing.T) {
	message := "my message"
	err := IsTrue(false, message)

	if err == nil {
		t.Errorf("IsTrue %s failed when expression is false", "case 1")
	}
	if _, ok := err.(*argument.Error); !ok {
		t.Errorf("IsTrue %s failed must return %s", "case 1", "Error")
	}

	if err.Error() != "my message" {
		t.Errorf("IsTrue %s failed to return message: %s", "case 1", message)
	}
}

func TestIsTrueCase2(t *testing.T) {
	err := IsTrue(false, "")

	if err == nil {
		t.Errorf("IsTrue %s failed when expression is false", "case 2")
	}
	if _, ok := err.(*argument.Error); !ok {
		t.Errorf("IsTrue %s failed must return %s", "case 2", "Error")
	}

	if err.Error() != validate.DefaultIsTrueExMessage {
		t.Errorf("IsTrue %s failed to return default message %s", "case 2", validate.DefaultIsTrueExMessage)
	}
}

func TestIsTrueCase3(t *testing.T) {
	err := IsTrue(true, "")
	if err != nil {
		t.Errorf("IsTrue failed to validate expression")
	}
}

func TestIsFalseCase1(t *testing.T) {
	message := "my message"
	err := IsFalse(true, message)

	if err == nil {
		t.Errorf("IsFalse %s failed when expression is true", "case 1")
	}
	if _, ok := err.(*argument.Error); !ok {
		t.Errorf("IsFalse %s failed must return %s", "case 1", "Error")
	}

	if err.Error() != "my message" {
		t.Errorf("IsFalse %s failed to return message: %s", "case 1", message)
	}
}

func TestIsFalseCase2(t *testing.T) {
	err := IsFalse(true, "")

	if err == nil {
		t.Errorf("IsFalse %s failed when expression is true", "case 2")
	}
	if _, ok := err.(*argument.Error); !ok {
		t.Errorf("IsFalse %s failed must return %s", "case 2", "Error")
	}

	if err.Error() != validate.DefaultIsTrueExMessage {
		t.Errorf("IsFalse %s failed to return default message %s", "case 2", validate.DefaultIsTrueExMessage)
	}
}

func TestIsFalseCase3(t *testing.T) {
	err := IsTrue(true, "")
	if err != nil {
		t.Errorf("IsFalse failed to validate expression")
	}
}
