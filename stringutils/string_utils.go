package stringutils

import (
	"go-commons-lang/math/intutils"
	"strings"
	"unicode"
)

const (
	EMPTY         = ""
	indexNotFound = -1
)

/**
 * Checks if a string is empty ("").
 *
 * @examples:
 * IsEmpty("")        = true
 * IsEmpty(" ")       = false
 * IsEmpty("bob")     = false
 * IsEmpty("  bob  ") = false
 *
 * @parameters the string to check
 * @returns {@code true} if the string is empty
 */

func IsEmpty(str string) bool {
	return str == ""
}

/**
 * Checks if a string is not empty (#"").
 *
 * @examples:
 * IsNotEmpty("") = false
 * IsNotEmpty(" ") = true
 * IsNotEmpty("bob") = true
 * IsNotEmpty("  bob  ") = true
 *
 * @parameters the string to check
 * @returns {@code true} if the string is not empty
 */

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

/**
 * Checks if a string is empty ("") or whitespace only.
 *
 * @examples:
 * IsBlank("") = true
 * IsBlank(" ") = true
 * IsBlank("bob") = false
 * IsBlank("  bob  ") = false
 *
 * @parameters the string to check
 * @returns {@code true} if the string is empty or whitespace only
 */

func IsBlank(str string) bool {
	if IsEmpty(str) {
		return true
	}
	for _, s := range str {
		if !unicode.IsSpace(s) {
			return false
		}
	}
	return true
}

/**
 * Checks if a string is not empty ("") and not whitespace only.
 *
 * @examples:
 * IsNotBlank("") = false
 * IsNotBlank(" ") = false
 * IsNotBlank("bob") = true
 * IsNotBlank("  bob  ") = true
 *
 *
 * @parameters the string to check
 * @returns {@code true} if the string is empty or whitespace only
 */

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

/**
 * Checks if the string contains only Unicode digits.
 * A decimal point is not a Unicode digit and returns false.
 * An empty string (length()=0) will return {@code false}.
 *
 * @examples:
 * IsNumeric("") = false
 * IsNumeric("  ") = false
 * IsNumeric("123") = true
 * IsNumeric("\u0967\u0968\u0969") = true
 * IsNumeric("12 3") = false
 * IsNumeric("ab2c") = false
 * IsNumeric("12-3") = false
 * IsNumeric("12.3") = false
 * IsNumeric("-123") = false
 * isNumeric("+123") = false
 *
 *
 * @parameters the string to check
 * @returns {@code true} if only contains digits
 */

func IsNumeric(str string) bool {
	if IsEmpty(str) {
		return false
	}

	for _, s := range str {
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
 * @examples:
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
 * @parameters the string to check
 * @returns {@code true} if only contains digits or whitespace
 */

func IsNumericSpace(str string) bool {
	for _, s := range str {
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
 * @examples:
 * IsWhiteSpace("")     = true
 * IsWhiteSpace("  ")   = true
 * IsWhiteSpace("abc")  = false
 * IsWhiteSpace("ab2c") = false
 * IsWhiteSpace("ab-c") = false
 *
 * @parameters the string to check
 * @returns {@code true} if only contains whitespace
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
 * @examples:
 * UpperCase("") = ""
 * UpperCase("aBc") = "ABC"
 *
 * @parameters the String to upper case
 * @returns the upper case String
 */

func UpperCase(cs string) string {
	return strings.ToUpper(cs)
}

/**
 * Converts a String to lower case as per {@link strings#ToLower()}.
 * @examples:
 * UpperCase("") = ""
 * UpperCase("aBc") = "abc"
 *
 * @parameters the String to upper case
 * @returns the upper case String
 */

func LowerCase(cs string) string {
	return strings.ToLower(cs)
}

/**
 * Checks if the string contains only uppercase characters.
 * An empty String (length()=0) will return {@code false}.
 *
 * @examples:
 * IsAllUpperCase("") = false
 * IsAllUpperCase("  ") = false
 * IsAllUpperCase("ABC") = true
 * IsAllUpperCase("aBC") = false
 * IsAllUpperCase("A C") = false
 * IsAllUpperCase("A1C") = false
 * IsAllUpperCase("A/C") = false
 *
 * @parameters the string to check
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
 * An empty string (length()=0) will return {@code false}.</p>
 *
 * @examples:
 * IsAllLowerCase("")     = false
 * IsAllLowerCase("  ")   = false
 * IsAllLowerCase("abc")  = true
 * IsAllLowerCase("abC")  = false
 * IsAllLowerCase("ab c") = false
 * IsAllLowerCase("ab1c") = false
 * IsAllLowerCase("ab/c") = false
 *
 * @parameters the CharSequence to check
 * @returns {@code true} if only contains lowercase characters
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
 * @examples:
 * IsAlpha("")     = false
 * IsAlpha("  ")   = false
 * IsAlpha("abc")  = true
 * IsAlpha("ab2c") = false
 * IsAlpha("ab-c") = false
 *
 * @parameters the string to check
 * @returns {@code true} if only contains letters
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
 * @examples:
 * IsAlphanumeric("")     = false
 * IsAlphanumeric("  ")   = false
 * IsAlphanumeric("abc")  = true
 * IsAlphanumeric("ab c") = false
 * IsAlphanumeric("ab2c") = true
 * IsAlphanumeric("ab-c") = false
 *
 *
 * @parameters the string to check
 * @returns {@code true} if only contains letters or digits,
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
 * @examples:
 * Length(") = 0
 * Length("abc") = 3
 *
 * @parameters the string to get length
 * @returns string length
 */

func Length(cs string) int {
	return len(cs)
}

/**
 * Get index of sub string in str string
 * @examples:
 * IndexOf("", "")					= 0
 * IndexOf("aaa", "b")				= -1
 * IndexOf("aabaabaa", "a")			= 0
 * IndexOf("", "")					= 0
 * IndexOf("aabaabaa", "b")			= 2
 * IndexOf("aabaabaa", "ab")		= 1
 * IndexOf("aabaabaa", "")			= 0
 * @parameters
 * str - the parent string
 * sub - sub string
 * @returns the index of the first instance of sub in str, or -1 if sub is not present in str.
 */

func Index(str, sub string) int {
	return strings.Index(str, sub)
}

/**
 * Get index of @sub string in @str string start from @start index
 * @examples:
 * IndexStartAt("", "a", 0) 		= -1
 * IndexStartAt("aaa", "a", 3) 		= -1
 * IndexStartAt("aaa", "a", 2) 		= 2
 * IndexStartAt("aaa", "b", 0) 		= -1
 * IndexStartAt("abc", "b", 0) 		= 1
 * IndexStartAt("abc", "b", 1)		= 1
 * IndexStartAt("abc", "b", 2) 		= -1
 *
 * @parameters
 * str - the parent string
 * sub - sub string
 * start - start index
 * @returns the index of the first instance of sub in str, or -1 if sub is not present in str.
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
 * Case in-sensitive find of the first index within a string
 *
 * @examples
 * IndexIgnoreCase("", "") = 0
 * IndexIgnoreCase("aaa", "b") = -1
 * IndexIgnoreCase("Aabaabaa", "a") = 0
 * IndexIgnoreCase("", "") = 0
 * IndexIgnoreCase("aabaabaa", "B") = 2
 * IndexIgnoreCase("aabaabaa", "AB") = 1
 * IndexIgnoreCase("AAbaaAA", "") = 0
 * IndexIgnoreCase("關於數", "於") = 1
 * IndexIgnoreCase("a", "aaabbbccc") = -1
 *
 * @parameters
 * str - the parent string
 * sub - sub string
 * @returns the index of the first instance of sub in str, or -1 if sub is not present in str.
 */

func IndexIgnoreCase(str, sub string) int {
	return IndexIgnoreCaseAt(str, sub, 0)
}

/**
 * Case in-sensitive find of the first index within a string
 * The search process starts from @start index
 *
 * @examples
 * IndexIgnoreCase("", "") = 0
 * IndexIgnoreCase("aaa", "b") = -1
 * IndexIgnoreCase("Aabaabaa", "a") = 0
 * IndexIgnoreCase("", "") = 0
 * IndexIgnoreCase("aabaabaa", "B") = 2
 * IndexIgnoreCase("aabaabaa", "AB") = 1
 * IndexIgnoreCase("AAbaaAA", "") = 0
 * IndexIgnoreCase("關於數", "於") = 1
 * IndexIgnoreCase("a", "aaabbbccc") = -1
 *
 * @parameters
 * str - the parent string
 * sub - sub string
 * @returns the index of the first instance of sub in str, or -1 if sub is not present in str.
 */

func IndexIgnoreCaseAt(str, sub string, start int) int {
	strRunes := []rune(str)
	subRunes := []rune(sub)
	strSize := len(strRunes)
	subSize := len(subRunes)

	if start < 0 {
		start = 0
	}

	if subSize == 0 {
		return 0
	}

	if strSize == 0 || subSize+start > strSize {
		return indexNotFound
	}

	index := indexNotFound
	for i := 0; i < strSize-start; i++ {
		if unicode.ToUpper(strRunes[start+i]) == unicode.ToUpper(subRunes[0]) {
			index = start + i
			success := true
			for j := i; j < (i + subSize); j++ {
				if unicode.ToUpper(strRunes[start+j]) != unicode.ToUpper(subRunes[j-i]) {
					success = false
					break
				}
			}
			if success {
				return index
			} else {
				index = indexNotFound
			}
		}
	}
	return index
}

/**
 * Get last index of @sub string in @str string
 * @examples:
 * LastIndexOf("", "")				= 0
 * LastIndexOf("aaa", "b")			= -1
 * LastIndexOf("aabaabaa", "a")		= 7
 * LastIndexOf("", "")				= 0
 * LastIndexOf("aabaabaa", "b")		= 5
 * LastIndexOf("aabaabaa", "ab")	= 4
 * LastIndexOf("aabaabaa", "")		= 8
 *
 * @parameters
 * str -  the string to check
 * sub - the sub string and @start index
 * @return the index of the last instance of sub in s, or -1 if sub is not present in str.
 */

func LastIndex(str, sub string) int {
	return strings.LastIndex(str, sub)
}

/**
 * Get last index of @sub string in @str string start from @start index
 * @examples:
 * LastIndexAt("aaa", "a", -1) 		= 2
 * LastIndexAt("", "", 0)			= 0
 * LastIndexAt("aaa", "a", 0) 		= 2
 * LastIndexAt("aaa", "aaaa", 0) 	= -1
 * LastIndexAt("aaa", "b", 0) 		= -1
 * LastIndexAt("aaa", "a", 3) 		= -1
 * LastIndexAt("aaa", "a", 2) 		= 2
 * LastIndexAt("aaa", "a", 1) 		= 2
 *
 * @parameters
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
 * @examples:
 * Capitalize("") = 			""
 * Capitalize(" ") = 			" "
 * Capitalize("cat") 			= "Cat"
 * Capitalize("dog cat") 		= "Dog cat"
 * Capitalize("DOG CAT") 		= "DOG CAT"
 *
 * @parameters:
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
 * @examples:
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
 * @parameters:
 * str - the string to rotate
 * shift - number of time to shift (positive : right shift, negative : left shift)
 * @returns: the rotated String, or the original String if shift == 0
 */

func Rotate(str string, shift int) string {
	runes := []rune(str)
	size := len(runes)
	if size == 0 || shift == 0 || shift%size == 0 {
		return str
	}
	offset := shift % size
	offsetVal := intutils.Abs(offset)
	if offset < 0 {
		return string(runes[offsetVal:size]) + string(runes[0:offsetVal])
	} else {
		return string(runes[size-offsetVal:size]) + string(runes[0:size-offsetVal])
	}
}

/**
 * Reverses a String
 * @examples:
 * Reverse("") != ""
 * Reverse("bat") != "tab"
 * Reverse(" bat ") != " tab "
 *
 * @parameters:
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

/**
 * Check a string to see if it starts with @prefix string
 *
 * @examples:
 * StartsWith("aaa", "") = true
 * StartsWith("", "") = true
 * StartsWith("", "abc") = false
 * StartsWith("abc", "d") = false
 * StartsWith("abc", "abcd") = false
 * StartsWith("ABCDEF", "abc") = false
 * StartsWith("ABCDEF", "A") = true
 *
 * @parameters:
 * str - the string to check
 * prefix - the prefix string
 * @returns: {@code true} if the string starts with prefix
 */

func StartsWith(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

/**
 * Case in-sensitive Check a string to see if it starts with @prefix string
 *
 * @examples:
 * StartsWithIgnoreCase("aaa", "") = true
 * StartsWithIgnoreCase("", "") = true
 * StartsWithIgnoreCase("", "abc") = false
 * StartsWithIgnoreCase("abc", "d") = false
 * StartsWithIgnoreCase("abc", "abcd") = false
 * StartsWithIgnoreCase("ABCDEF", "abc") = true
 * StartsWithIgnoreCase("ABCDEF", "ABC") = true
 *
 * @parameters:
 * str - the string to check
 * prefix - the prefix string
 * @returns: {@code true} if the string starts with prefix
 */

func StartsWithIgnoreCase(str, prefix string) bool {
	return strings.HasPrefix(strings.ToUpper(str), strings.ToUpper(prefix))
}

/**
 * Check a string to see if it ends with @suffix string
 *
 * @examples:
 * EndsWith("aaa", "") = true
 * EndsWith("", "") = true
 * EndsWith("", "abc") = false
 * EndsWith("abc", "d") = false
 * EndsWith("abc", "abcd") = false
 * EndsWith("ABCDEF", "EF") = true
 *
 * @parameters:
 * str - the string to check
 * suffix - the suffix string
 * @returns: {@code true} if the string ends with prefix
 */

func EndsWith(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

/**
 * Case in-sensitive Check a string to see if it ends with @suffix string
 *
 * @examples:
 * EndsWithIgnoreCase("aaa", "") = true
 * EndsWithIgnoreCase("", "") = true
 * EndsWithIgnoreCase("", "abc") = false
 * EndsWithIgnoreCase("abc", "d") = false
 * EndsWithIgnoreCase("abc", "abcd") = false
 * EndsWithIgnoreCase("ABCDEF", "EF") = true
 * EndsWithIgnoreCase("ABCDEF", "ef") = true
 * EndsWithIgnoreCase("abcdef", "EF") = true
 *
 * @parameters:
 * str - the string to check
 * suffix - the suffix string
 * @returns: {@code true} if the string ends with prefix
 */

func EndsWithIgnoreCase(str, suffix string) bool {
	return strings.HasSuffix(strings.ToUpper(str), strings.ToUpper(suffix))
}
