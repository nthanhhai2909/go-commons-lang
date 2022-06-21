package stringutils

import (
	"strings"
	"unicode"
)

const (
	EMPTY = ""
)

/**
 * <p>Checks if a string is empty ("").</p>
 *
 * <pre>
 * IsEmpty("")        = true
 * IsEmpty(" ")       = false
 * IsEmpty("bob")     = false
 * IsEmpty("  bob  ") = false
 * </pre>
 *
 * @param cs the string to check
 * @return {@code true} if the string is empty
 */

func IsEmpty(cs string) bool {
	return cs == ""
}

/**
 * <p>Checks if a string is not empty (#"").</p>
 *
 * <pre>
 * IsNotEmpty("")        = false
 * IsNotEmpty(" ")       = true
 * IsNotEmpty("bob")     = true
 * IsNotEmpty("  bob  ") = true
 * </pre>
 *
 * @param cs the string to check
 * @return {@code true} if the string is not empty
 */

func IsNotEmpty(cs string) bool {
	return !IsEmpty(cs)
}

/**
 * <p>Checks if a string is empty ("") or whitespace only.</p>
 *
 * <pre>
 * IsBlank("")        = true
 * IsBlank(" ")       = true
 * IsBlank("bob")     = false
 * IsBlank("  bob  ") = false
 * </pre>
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
 * <p>Checks if a string is not empty ("") and not whitespace only.</p>
 *
 * <pre>
 * IsNotBlank("")        = false
 * IsNotBlank(" ")       = false
 * IsNotBlank("bob")     = true
 * IsNotBlank("  bob  ") = true
 * </pre>
 *
 * @param cs the string to check
 * @return {@code true} if the string is empty or whitespace only
 */

func IsNotBlank(cs string) bool {
	return !IsBlank(cs)
}

/**
 * <p>Checks if the string contains only Unicode digits.
 * A decimal point is not a Unicode digit and returns false.</p>
 * An empty string (length()=0) will return {@code false}.</p>
 *
 * <pre>
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
 * </pre>
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
 * <p>Checks if the string contains only Unicode digits or whitespace.
 * A decimal point is not a Unicode digit and returns false.</p>
 * An empty string (length()=0) will return {@code false}.</p>
 *
 * <pre>
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
 * </pre>
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
 * <p>Checks if the string contains only whitespace.</p>
 *
 * An empty string (length()=0) will return {@code true}.</p>
 *
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
 * <p>Converts a String to upper case as per {@link strings#ToUpper()}.</p>
 * <pre>
 * UpperCase("")    = ""
 * UpperCase("aBc") = "ABC"
 * </pre>
 * @param str  the String to upper case
 * @return the upper case String
 */

func UpperCase(cs string) string {
	return strings.ToUpper(cs)
}

/**
 * <p>Converts a String to lowser case as per {@link strings#ToLower()}.</p>
 * <pre>
 * UpperCase("")    = ""
 * UpperCase("aBc") = "abc"
 * </pre>
 * @param str  the String to upper case
 * @return the upper case String
 */

func LowerCase(cs string) string {
	return strings.ToLower(cs)
}

/**
 * <p>Checks if the string contains only uppercase characters.</p>
 *
 * An empty String (length()=0) will return {@code false}.</p>
 *
 * <pre>
 * IsAllUpperCase("")     = false
 * IsAllUpperCase("  ")   = false
 * IsAllUpperCase("ABC")  = true
 * IsAllUpperCase("aBC")  = false
 * IsAllUpperCase("A C")  = false
 * IsAllUpperCase("A1C")  = false
 * IsAllUpperCase("A/C")  = false
 * </pre>
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
 * <p>Checks if the string contains only lowercase characters.</p>
 *
 * An empty string (length()=0) will return {@code false}.</p>
 *
 * <pre>
 * IsAllLowerCase("")     = false
 * IsAllLowerCase("  ")   = false
 * IsAllLowerCase("abc")  = true
 * IsAllLowerCase("abC")  = false
 * IsAllLowerCase("ab c") = false
 * IsAllLowerCase("ab1c") = false
 * IsAllLowerCase("ab/c") = false
 * </pre>
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
 * <p>Checks if the string contains only Unicode letters.</p>
 *
 * An empty CharSequence (length()=0) will return {@code false}.</p>
 *
 * <pre>
 * IsAlpha("")     = false
 * IsAlpha("  ")   = false
 * IsAlpha("abc")  = true
 * IsAlpha("ab2c") = false
 * IsAlpha("ab-c") = false
 * </pre>
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
 * <p>Checks if the string contains only Unicode letters or digits.</p>
 *
 * An empty CharSequence (length()=0) will return {@code false}.</p>
 *
 * <pre>
 * IsAlphanumeric("")     = false
 * IsAlphanumeric("  ")   = false
 * IsAlphanumeric("abc")  = true
 * IsAlphanumeric("ab c") = false
 * IsAlphanumeric("ab2c") = true
 * IsAlphanumeric("ab-c") = false
 * </pre>
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
 * <p>Get the string length.</p>
 * <pre>
 * Length(") = 0
 * Length("abc") = 3
 * </pre>
 * @param cs  the string to get length
 * @return string length
 */

func Length(cs string) int {
	return len(cs)
}

/**
 * @param @str string and @sub substr to check index
 * <pre>
 * IndexOf("", "") 					= 0
 * IndexOf("aaa", "b") 				= -1
 * IndexOf("aabaabaa", "a") 		= 0
 * IndexOf("", "") 					= 0
 * IndexOf("aabaabaa", "b") 		= 2
 * IndexOf("aabaabaa", "ab") 		= 1
 * IndexOf("aabaabaa", "") 			= 0
 * </pre>
 * @return the index of the first instance of substr in s, or -1 if substr is not present in str.
 */

func IndexOf(str, sub string) int {
	return strings.Index(str, sub)
}

/**
 * @param @str string and @sub substr to check index
 * <pre>
 * LastIndexOf("", "")				= 0
 * LastIndexOf("aaa", "b")  		= -1
 * LastIndexOf("aabaabaa", "a") 	= 7
 * LastIndexOf("", "") 				= 0
 * LastIndexOf("aabaabaa", "b") 	= 5
 * LastIndexOf("aabaabaa", "ab") 	= 4
 * LastIndexOf("aabaabaa", "") 		= 8
 * </pre>
 * @return the index of the last instance of substr in s, or -1 if substr is not present in str.
 */

func LastIndexOf(str, sub string) int {
	return strings.LastIndex(str, sub)
}

/**
 * Capitalizes a String changing the first character to title case as
 * <pre>
 * Capitalize("") = 			""
 * Capitalize(" ") = 			" "
 * Capitalize("cat") 			= "Cat"
 * Capitalize("dog cat") 		= "Dog cat"
 * Capitalize("DOG CAT") 		= "DOG CAT"
 * </pre>
 * @param the string to capitalize
 * @return the capitalized String
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
