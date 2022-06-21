package stringutils

import "testing"

func TestIsEmpty(t *testing.T) {
	if !IsEmpty(EMPTY) ||
		IsEmpty(" ") ||
		IsEmpty("bob") ||
		IsEmpty("  bob  ") {
		t.Errorf("IsEmpty func don't work correctly")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if IsNotEmpty(EMPTY) ||
		!IsNotEmpty(" ") ||
		!IsNotEmpty("bob") ||
		!IsNotEmpty("  bob  ") {
		t.Errorf("IsNotEmpty func don't work correctly")
	}
}

func TestIsBlank(t *testing.T) {
	if !IsBlank(EMPTY) ||
		!IsBlank(" ") ||
		IsBlank("bob") ||
		IsBlank("  bob  ") {
		t.Errorf("IsBlank func don't work correctly")
	}
}

func TestIsNotBlank(t *testing.T) {
	if IsNotBlank(EMPTY) ||
		IsNotBlank(" ") ||
		!IsNotBlank("bob") ||
		!IsNotBlank("  bob  ") {
		t.Errorf("IsNotBlank func don't work correctly")
	}
}

func TestIsNumeric(t *testing.T) {
	if IsNumeric(EMPTY) ||
		IsNumeric("  ") ||
		IsNumeric("12 3") ||
		IsNumeric("ab2c") ||
		IsNumeric("12-3") ||
		IsNumeric("12.3") ||
		IsNumeric("-123") ||
		IsNumeric("+123") ||
		!IsNumeric("123") ||
		!IsNumeric("\u0967\u0968\u0969") {
		t.Errorf("IsNumeric func don't work correctly")
	}
}

func TestIsNumericSpace(t *testing.T) {
	if !IsNumericSpace("  ") ||
		!IsNumericSpace("123") ||
		!IsNumericSpace("\u0967\u0968\u0969") ||
		!IsNumericSpace("12 3") ||
		!IsNumericSpace("") ||
		IsNumericSpace("ab2c") ||
		IsNumericSpace("12-3") ||
		IsNumericSpace("12.3") ||
		IsNumericSpace("-123") ||
		IsNumericSpace("+123") {
		t.Errorf("IsNumericSpace func don't work correctly")
	}
}

func TestIsWhiteSpace(t *testing.T) {
	if !IsWhiteSpace(EMPTY) ||
		!IsWhiteSpace("  ") ||
		IsWhiteSpace("abc") ||
		IsWhiteSpace("ab2c") ||
		IsWhiteSpace("ab-c") {
		t.Errorf("IsWhiteSpace func don't work correctly")
	}
}

func TestUpperCase(t *testing.T) {
	if UpperCase(EMPTY) != "" ||
		UpperCase("aBc") != "ABC" ||
		UpperCase("abc") != "ABC" {
		t.Errorf("UpperCase func don't work correctly")
	}
}

func TestLowerCase(t *testing.T) {
	if LowerCase(EMPTY) != "" ||
		LowerCase("aBc") != "abc" ||
		LowerCase("abc") != "abc" ||
		LowerCase("ABC") != "abc" {
		t.Errorf("LowerCase func don't work correctly")
	}
}

func TestIsAllUpperCase(t *testing.T) {
	if IsAllUpperCase(EMPTY) ||
		IsAllUpperCase("  ") ||
		IsAllUpperCase("aBC") ||
		IsAllUpperCase("A C") ||
		IsAllUpperCase("A1C") ||
		IsAllUpperCase("A/C") ||
		!IsAllUpperCase("ABC") {
		t.Errorf("IsAllUpperCase func don't work correctly")
	}
}

func TestIsAllLowerCase(t *testing.T) {
	if IsAllLowerCase(EMPTY) ||
		IsAllLowerCase("  ") ||
		IsAllLowerCase("abC") ||
		IsAllLowerCase("ab c") ||
		IsAllLowerCase("ab1c") ||
		IsAllLowerCase("ab/c") ||
		!IsAllLowerCase("abc") {
		t.Errorf("IsAllLowerCase func don't work correctly")
	}
}

func TestIsAlpha(t *testing.T) {
	if IsAlpha(EMPTY) ||
		IsAlpha("  ") ||
		!IsAlpha("abc") ||
		IsAlpha("ab2c") ||
		IsAlpha("ab-c") {
		t.Errorf("IsAlpha func don't work correctly")
	}
}

func TestIsAlphanumeric(t *testing.T) {
	if IsAlphanumeric(EMPTY) ||
		IsAlphanumeric("  ") ||
		!IsAlphanumeric("abc") ||
		IsAlphanumeric("ab c") ||
		!IsAlphanumeric("ab2c") ||
		IsAlphanumeric("ab-c") {
		t.Errorf("IsAlphanumeric func don't work correctly")
	}
}
