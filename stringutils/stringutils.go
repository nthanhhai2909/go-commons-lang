package stringutils

import "unicode"

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
