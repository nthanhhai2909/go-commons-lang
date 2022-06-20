package stringutils

import "testing"

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") ||
		IsEmpty(" ") ||
		IsEmpty("bob") ||
		IsEmpty("  bob  ") {
		t.Errorf("IsEmpty func don't work correctly")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty("") ||
		!IsNotEmpty(" ") ||
		!IsNotEmpty("bob") ||
		!IsNotEmpty("  bob  ") {
		t.Errorf("IsNotEmpty func don't work correctly")
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank("") ||
		!IsBlank(" ") ||
		IsBlank("bob") ||
		IsBlank("  bob  ") {
		t.Errorf("IsBlank func don't work correctly")
	}
}

func TestIsNotBlank(t *testing.T) {
	if IsNotBlank("") ||
		IsNotBlank(" ") ||
		!IsNotBlank("bob") ||
		!IsNotBlank("  bob  ") {
		t.Errorf("IsNotBlank func don't work correctly")
	}
}
