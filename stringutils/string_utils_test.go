package stringutils

import (
	"testing"
)

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

func TestLength(t *testing.T) {
	if Length(EMPTY) != 0 || Length("abc") != 3 {
		t.Errorf("Length func don't work correctly")
	}
}

func TestIndex(t *testing.T) {
	if Index("", "") != 0 ||
		Index("aaa", "b") != -1 ||
		Index("aabaabaa", "a") != 0 ||
		Index("", "") != 0 ||
		Index("aabaabaa", "b") != 2 ||
		Index("aabaabaa", "ab") != 1 ||
		Index("aabaabaa", "") != 0 {
		t.Errorf("Index func don't work correctly")
	}
}

func TestIndexStartAt(t *testing.T) {
	if IndexAt("aaa", "a", -1) != 0 ||
		IndexAt("", "a", 0) != -1 ||
		IndexAt("aaa", "a", 3) != -1 ||
		IndexAt("aaa", "a", 2) != 2 ||
		IndexAt("aaa", "b", 0) != -1 ||
		IndexAt("abc", "b", 0) != 1 ||
		IndexAt("abc", "b", 1) != 1 ||
		IndexAt("abc", "b", 2) != -1 {
		t.Errorf("IndexStartAt func don't work correctly")
	}
}

func TestIndexIgnoreCase(t *testing.T) {
	if IndexIgnoreCase("", "") != 0 ||
		IndexIgnoreCase("aaa", "b") != -1 ||
		IndexIgnoreCase("Aabaabaa", "a") != 0 ||
		IndexIgnoreCase("", "") != 0 ||
		IndexIgnoreCase("aabaabaa", "B") != 2 ||
		IndexIgnoreCase("aabaabaa", "AB") != 1 ||
		IndexIgnoreCase("AAbaaAA", "") != 0 ||
		IndexIgnoreCase("關於數", "於") != 1 ||
		IndexIgnoreCase("a", "aaabbbccc") != -1 {
		t.Errorf("IndexIgnoreCase func don't work correctly")
	}
}

func TestIndexIgnoreCaseAt(t *testing.T) {
	if IndexIgnoreCaseAt("", "", 0) != 0 ||
		IndexIgnoreCaseAt("", "", 1) != 0 ||
		IndexIgnoreCaseAt("aaa", "b", 0) != -1 ||
		IndexIgnoreCaseAt("Aabaabaa", "a", 1) != 1 ||
		IndexIgnoreCaseAt("aabaabaa", "B", 3) != 5 ||
		IndexIgnoreCaseAt("aabaabaa", "AB", 0) != 1 ||
		IndexIgnoreCaseAt("AAbaaAA", "", 9) != 0 ||
		IndexIgnoreCaseAt("關於數於", "於", 0) != 1 ||
		IndexIgnoreCaseAt("關於數於", "於", 2) != 3 ||
		IndexIgnoreCaseAt("關於數於", "於", 3) != 3 ||
		IndexIgnoreCaseAt("關於數於", "於", 4) != -1 ||
		IndexIgnoreCaseAt("a", "aaabbbccc", 0) != -1 ||
		IndexIgnoreCaseAt("AAbaaAA", "a", -10) != 0 {
		t.Errorf("IndexIgnoreCaseAt func don't work correctly")
	}
}

func TestLastIndex(t *testing.T) {
	if LastIndex("", "") != 0 ||
		LastIndex("aaa", "b") != -1 ||
		LastIndex("aabaabaa", "a") != 7 ||
		LastIndex("", "") != 0 ||
		LastIndex("aabaabaa", "b") != 5 ||
		LastIndex("aabaabaa", "ab") != 4 ||
		LastIndex("aabaabaa", "") != 8 {
		t.Errorf("LastIndex func don't work correctly")
	}
}

func TestLastIndexStartAt(t *testing.T) {
	if LastIndexAt("aaa", "a", -1) != 2 ||
		LastIndexAt("", "", 0) != 0 ||
		LastIndexAt("aaa", "a", 0) != 2 ||
		LastIndexAt("aaa", "aaaa", 0) != -1 ||
		LastIndexAt("aaa", "b", 0) != -1 ||
		LastIndexAt("aaa", "a", 3) != -1 ||
		LastIndexAt("aaa", "a", 2) != 2 ||
		LastIndexAt("aaa", "a", 1) != 2 {
		t.Errorf("LastIndexAt func don't work correctly")
	}
}

func TestCapitalize(t *testing.T) {
	if Capitalize("") != "" ||
		Capitalize(" ") != " " ||
		Capitalize("cat") != "Cat" ||
		Capitalize("dog cat") != "Dog cat" ||
		Capitalize("DOG CAT") != "DOG CAT" {
		t.Errorf("Capitalize func don't work correctly")
	}
}

func TestRotate(t *testing.T) {
	if Rotate("", 0) != "" ||
		Rotate("", -2) != "" ||
		Rotate("", 2) != "" ||
		Rotate("abcdefg", 2) != "fgabcde" ||
		Rotate("abcdefg", -2) != "cdefgab" ||
		Rotate("abcdefg", 7) != "abcdefg" ||
		Rotate("abcdefg", -7) != "abcdefg" ||
		Rotate("abcdefg", 9) != "fgabcde" ||
		Rotate("abcdefg", -9) != "cdefgab" ||
		Rotate("關於數", 1) != "數關於" {
		t.Errorf("Rotate func don't work correctly")
	}
}

func TestReverse(t *testing.T) {
	if Reverse("") != "" ||
		Reverse("bat") != "tab" ||
		Reverse(" bat ") != " tab " ||
		Reverse("關於數") != "數於關" {
		t.Errorf("Reverse func don't work correctly")
	}
}

func TestStartsWith(t *testing.T) {
	if !StartsWith("aaa", "") ||
		!StartsWith("", "") ||
		StartsWith("", "abc") ||
		StartsWith("abc", "d") ||
		StartsWith("abc", "abcd") ||
		StartsWith("ABCDEF", "abc") ||
		!StartsWith("ABCDEF", "A") {
		t.Errorf("StartsWith func don't work correctly")
	}
}

func TestStartsWithIgnoreCase(t *testing.T) {
	if !StartsWithIgnoreCase("aaa", "") ||
		!StartsWithIgnoreCase("", "") ||
		StartsWithIgnoreCase("", "abc") ||
		StartsWithIgnoreCase("abc", "d") ||
		StartsWithIgnoreCase("abc", "abcd") ||
		!StartsWithIgnoreCase("ABCDEF", "abc") ||
		!StartsWithIgnoreCase("ABCDEF", "ABC") {
		t.Errorf("StartsWithIgnoreCase func don't work correctly")
	}
}

func TestEndsWith(t *testing.T) {
	if !EndsWith("aaa", "") ||
		!EndsWith("", "") ||
		EndsWith("", "abc") ||
		EndsWith("abc", "d") ||
		EndsWith("abc", "abcd") ||
		!EndsWith("ABCDEF", "EF") {
		t.Errorf("EndsWith func don't work correctly")
	}
}

func TestEndsWithIgnoreCase(t *testing.T) {
	if !EndsWithIgnoreCase("aaa", "") ||
		!EndsWithIgnoreCase("", "") ||
		EndsWithIgnoreCase("", "abc") ||
		EndsWithIgnoreCase("abc", "d") ||
		EndsWithIgnoreCase("abc", "abcd") ||
		!EndsWithIgnoreCase("ABCDEF", "EF") ||
		!EndsWithIgnoreCase("ABCDEF", "ef") ||
		!EndsWithIgnoreCase("abcdef", "EF") {
		t.Errorf("TestEndsWithIgnoreCase func don't work correctly")
	}
}
