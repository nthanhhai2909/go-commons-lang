package boolean

import (
	"go-commons-lang/errors"
	"go-commons-lang/errors/illegal/argument"
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

	if err.Error() != errors.DefaultIsTrueExMessage {
		t.Errorf("IsTrue %s failed to return default message %s", "case 2", errors.DefaultIsTrueExMessage)
	}
}

func TestIsTrueCase3(t *testing.T) {
	err := IsTrue(true, "")
	if err != nil {
		t.Errorf("IsTrue failed to validate expression")
	}
}
