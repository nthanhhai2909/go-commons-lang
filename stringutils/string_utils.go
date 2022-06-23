package stringutils

import (
	"go-commons-lang/math/intutils"
	"strings"
	"unicode"
)

const (
	EMPTY = ""
)

/**
 * Checks if a string is empty ("").
 *
 * @Examples:
 * IsEmpty("")        = true
 * IsEmpty(" ")       = false
 * IsEmpty("bob")     = false
 * IsEmpty("  bob  ") = false
 *
 *
 * @param cs the string to check
 * @return {@code true} if the string is empty
 */

func IsEmpty(cs string) bool {
	return cs == ""
}

/**
 * Checks if a string is not empty (#"").
 *
 * @Examples:
 * IsNotEmpty("")        = false
 * IsNotEmpty(" ")       = true
 * IsNotEmpty("bob")     = true
 * IsNotEmpty("  bob  ") = true
 *
 *
 * @param cs the string to check
 * @return {@code true} if the string is not empty
 */

func IsNotEmpty(cs string) bool {
	return !IsEmpty(cs)
}

/**
 * Checks if a string is empty ("") or whitespace only.
 *
 * @Examples:
 * IsBlank("")        = true
 * IsBlank(" ")       = true
 * IsBlank("bob")     = false
 * IsBlank("  bob  ") = false
 *
 *
 * @param cs the string to check
 * @return {@code true} if the string is empty or whitespace only
 */

func IsBlank(cs string) bool {
	if IsEmpty(cs) {
		return true
	}
	for _, s := range cs {
		if !unicode.IsSpace(s) {
			return false
		}
	}
	return true
}

/**
 * Checks if a string is not empty ("") and not whitespace only.
 *
 * @Examples:
 * IsNotBlank("")        = false
 * IsNotBlank(" ")       = false
 * IsNotBlank("bob")     = true
 * IsNotBlank("  bob  ") = true
 *
 *
 * @param cs the string to check
 * @return {@code true} if the string is empty or whitespace only
 */

func IsNotBlank(cs string) bool {
	return !IsBlank(cs)
}

/**
 * Checks if the string contains only Unicode digits.
 * A decimal point is not a Unicode digit and returns false.
 * An empty string (length()=0) will return {@code false}.
 *
 * @Examples:
 * IsNumeric("")     = false
 * IsNumeric("  ")   = false
 * IsNumeric("123")  = true
 * IsNumeric("\u0967\u0968\u0969")  = true
 * IsNumeric("12 3") = false
 * IsNumeric("ab2c") = false
 * IsNumeric("12-3") = false
 * IsNumeric("12.3") = false
 * IsNumeric("-123") = false
 * isNumeric("+123") = false
 *
 *
 * @param cs  the string to check
 * @return {@code true} if only contains digits
 */

func IsNumeric(cs string) bool {
	if IsEmpty(cs) {
		return false
	}

	for _, s := range cs {
		if !unicode.IsDigit(s) {
			return false
		}
	}
	return true
}

/**
 * Checks if the string contains only Unicode digits or whitespace.
 * A decimal point is not a Unicode digit and returns false.
 * An empty string (length()=0) will return {@code false}.
 *
 * @Examples:
 * IsNumericSpace("")     = true
 * IsNumericSpace("  ")   = true
 * IsNumericSpace("123")  = true
 * IsNumericSpace("\u0967\u0968\u0969")  = true
 * IsNumericSpace("12 3") = true
 * IsNumericSpace("ab2c") = false
 * IsNumericSpace("12-3") = false
 * IsNumericSpace("12.3") = false
 * IsNumericSpace("-123") = false
 * IsNumericSpace("+123") = false
 *
 *
 * @param cs  the string to check
 * @return {@code true} if only contains digits
 */

func IsNumericSpace(cs string) bool {
	for _, s := range cs {
		if !unicode.IsSpace(s) && !unicode.IsDigit(s) {
			return false
		}
	}
	return true
}

/**
 * Checks if the string contains only whitespace.
 *
 * An empty string (length()=0) will return {@code true}.
 *
 * @Examples:
 * IsWhiteSpace("")     = true
 * IsWhiteSpace("  ")   = true
 * IsWhiteSpace("abc")  = false
 * IsWhiteSpace("ab2c") = false
 * IsWhiteSpace("ab-c") = false
 * </pre>
 *
 * @param cs  the CharSequence to check
 * @return {@code true} if only contains whitespace
 */

func IsWhiteSpace(cs string) bool {
	for _, s := range cs {
		if !unicode.IsSpace(s) {
			return false
		}
	}

	return true
}

/**
 * Converts a String to upper case as per {@link strings#ToUpper()}.
 * @Examples:
 * UpperCase("")    = ""
 * UpperCase("aBc") = "ABC"
 *
 * @param str  the String to upper case
 * @return the upper case String
 */

func UpperCase(cs string) string {
	return strings.ToUpper(cs)
}

/**
 * Converts a String to lowser case as per {@link strings#ToLower()}.
 * @Examples:
 * UpperCase("")    = ""
 * UpperCase("aBc") = "abc"
 *
 * @param str  the String to upper case
 * @return the upper case String
 */

func LowerCase(cs string) string {
	return strings.ToLower(cs)
}

/**
 * Checks if the string contains only uppercase characters.
 * An empty String (length()=0) will return {@code false}.
 *
 * @Examples:
 * IsAllUpperCase("")     = false
 * IsAllUpperCase("  ")   = false
 * IsAllUpperCase("ABC")  = true
 * IsAllUpperCase("aBC")  = false
 * IsAllUpperCase("A C")  = false
 * IsAllUpperCase("A1C")  = false
 * IsAllUpperCase("A/C")  = false
 *
 *
 * @param cs the string to check
 * @return {@code true} if only contains uppercase characters
 */

func IsAllUpperCase(cs string) bool {
	if IsEmpty(cs) {
		return false
	}
	for _, s := range cs {
		if !unicode.IsUpper(s) {
			return false
		}
	}
	return true
}

/**
 * Checks if the string contains only lowercase characters
 *
 * An empty string (length()=0) will return {@code false}.</p>
 *
 * @Examples:
 * IsAllLowerCase("")     = false
 * IsAllLowerCase("  ")   = false
 * IsAllLowerCase("abc")  = true
 * IsAllLowerCase("abC")  = false
 * IsAllLowerCase("ab c") = false
 * IsAllLowerCase("ab1c") = false
 * IsAllLowerCase("ab/c") = false
 *
 *
 * @param cs  the CharSequence to check
 * @return {@code true} if only contains lowercase characters
 */

func IsAllLowerCase(cs string) bool {
	if IsEmpty(cs) {
		return false
	}

	for _, s := range cs {
		if !unicode.IsLower(s) {
			return false
		}
	}

	return true
}

/**
 * Checks if the string contains only Unicode letters.
 * An empty string (length()=0) will return {@code false}.
 * @Examples:
 * IsAlpha("")     = false
 * IsAlpha("  ")   = false
 * IsAlpha("abc")  = true
 * IsAlpha("ab2c") = false
 * IsAlpha("ab-c") = false
 *
 *
 * @param cs  the string to check
 * @return {@code true} if only contains letters
 */

func IsAlpha(cs string) bool {
	if IsEmpty(cs) {
		return false
	}

	for _, s := range cs {
		if !unicode.IsLetter(s) {
			return false
		}
	}

	return true
}

/**
 * Checks if the string contains only Unicode letters or digits.
 *
 * @Examples:
 * IsAlphanumeric("")     = false
 * IsAlphanumeric("  ")   = false
 * IsAlphanumeric("abc")  = true
 * IsAlphanumeric("ab c") = false
 * IsAlphanumeric("ab2c") = true
 * IsAlphanumeric("ab-c") = false
 *
 *
 * @param cs  the CharSequence to check
 * @return {@code true} if only contains letters or digits,
 */

func IsAlphanumeric(cs string) bool {
	if IsEmpty(cs) {
		return false
	}

	for _, s := range cs {
		if !unicode.IsLetter(s) && !unicode.IsDigit(s) {
			return false
		}
	}

	return true
}

/**
 * Get the string length.
 * @Examples:
 * Length(") = 0
 * Length("abc") = 3
 *
 * @param cs  the string to get length
 * @return string length
 */

func Length(cs string) int {
	return len(cs)
}

/**
 * Get index of sub string in str string
 * @Examples:
 * IndexOf("", "")					= 0
 * IndexOf("aaa", "b")				= -1
 * IndexOf("aabaabaa", "a")			= 0
 * IndexOf("", "")					= 0
 * IndexOf("aabaabaa", "b")			= 2
 * IndexOf("aabaabaa", "ab")		= 1
 * IndexOf("aabaabaa", "")			= 0
 * @Parameters
 * str - the parent string
 * sub - sub string
 * @return the index of the first instance of sub in str, or -1 if sub is not present in str.
 */

func Index(str, sub string) int {
	return strings.Index(str, sub)
}

/**
 * Get index of @sub string in @str string start from @start index
 * @Examples:
 * IndexStartAt("", "a", 0) 		= -1
 * IndexStartAt("aaa", "a", 3) 		= -1
 * IndexStartAt("aaa", "a", 2) 		= 2
 * IndexStartAt("aaa", "b", 0) 		= -1
 * IndexStartAt("abc", "b", 0) 		= 1
 * IndexStartAt("abc", "b", 1)		= 1
 * IndexStartAt("abc", "b", 2) 		= -1
 *
 * @Parameters
 * str - the parent string
 * sub - sub string
 * start - start index
 * @return the index of the first instance of sub in str, or -1 if sub is not present in str.
 */

func IndexAt(str, sub string, start int) int {
	if start < 0 {
		start = 0
	}

	if start > len(str) {
		return -1
	}
	index := strings.Index(str[start:], sub)
	if index == -1 {
		return index
	}
	return index + start
}

/**
 * Get last index of @sub string in @str string
 * @Examples:
 * LastIndexOf("", "")				= 0
 * LastIndexOf("aaa", "b")			= -1
 * LastIndexOf("aabaabaa", "a")		= 7
 * LastIndexOf("", "")				= 0
 * LastIndexOf("aabaabaa", "b")		= 5
 * LastIndexOf("aabaabaa", "ab")	= 4
 * LastIndexOf("aabaabaa", "")		= 8
 *
 * @Parameters
 * str -  the string to check
 * sub - the sub string and @start index
 * @return the index of the last instance of sub in s, or -1 if sub is not present in str.
 */

func LastIndex(str, sub string) int {
	return strings.LastIndex(str, sub)
}

/**
 * Get last index of @sub string in @str string start from @start index
 * @Examples:
 * LastIndexAt("aaa", "a", -1) 		= 2
 * LastIndexAt("", "", 0)			= 0
 * LastIndexAt("aaa", "a", 0) 		= 2
 * LastIndexAt("aaa", "aaaa", 0) 	= -1
 * LastIndexAt("aaa", "b", 0) 		= -1
 * LastIndexAt("aaa", "a", 3) 		= -1
 * LastIndexAt("aaa", "a", 2) 		= 2
 * LastIndexAt("aaa", "a", 1) 		= 2
 *
 * @Parameters
 * str -  the string to check
 * sub - the sub string and
 * start - start index
 * @return the index of the first instance of sub in str, or -1 if sub is not present in str.
 */

func LastIndexAt(str, sub string, start int) int {
	if start < 0 {
		start = 0
	}

	if start > len(str) {
		return -1
	}
	index := strings.LastIndex(str[start:], sub)
	if index == -1 {
		return index
	}
	return index + start
}

/**
 * Capitalizes a String changing the first character to title case as
 * @Examples:
 * Capitalize("") = 			""
 * Capitalize(" ") = 			" "
 * Capitalize("cat") 			= "Cat"
 * Capitalize("dog cat") 		= "Dog cat"
 * Capitalize("DOG CAT") 		= "DOG CAT"
 *
 * @Parameters:
 * cs - the string to capitalize
 * @returns: the capitalized String
 */

func Capitalize(cs string) string {
	if IsEmpty(cs) {
		return cs
	}

	runes := make([]rune, 0)
	index := 0
	for _, c := range cs {
		runes = append(runes, c)
		index++
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

/**
 * Rotate (circular shift) a String of shift characters.
 * @Examples:
 * Rotate("", 0)				= ""
 * Rotate("", -2)				= ""
 * Rotate("", 2)				= ""
 * Rotate("abcdefg", 2)			= "fgabcde"
 * Rotate("abcdefg", -2)		= "cdefgab"
 * Rotate("abcdefg", 7)			= "abcdefg"
 * Rotate("abcdefg", -7)		= "abcdefg"
 * Rotate("abcdefg", 9)			= "fgabcde"
 * Rotate("abcdefg", -9)		= "cdefgab"
 *
 * @Parameters:
 * str - the string to rotate
 * shift - number of time to shift (positive : right shift, negative : left shift)
 * @returns: the rotated String, or the original String if shift == 0
 */

func Rotate(str string, shift int) string {
	size := Length(str)
	if size == 0 || shift == 0 || shift%size == 0 {
		return str
	}
	offset := shift % size
	offsetVal := intutils.Abs(offset)
	if offset < 0 {
		return str[offsetVal:size] + str[0:offsetVal]
	} else {
		return str[size-offsetVal:size] + str[0:size-offsetVal]
	}
}

/**
 * Reverses a String
 * @Examples:
 * Reverse("") != ""
 * Reverse("bat") != "tab"
 * Reverse(" bat ") != " tab "
 *
 * @Parameters:
 * str - the string to reverse
 * @returns: the reversed String
 */

func Reverse(str string) string {
	runes := []rune(str)
	size := len(runes)
	for i := 0; i < size/2; i++ {
		swap := runes[i]
		runes[i] = runes[size-i-1]
		runes[size-i-1] = swap
	}
	return string(runes)
}