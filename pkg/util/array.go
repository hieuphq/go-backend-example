package util

import "strings"

// RemoveFirstElementBySeparator remove first elem in string with separator
// key: string - "user.name.id"
// separator: string - "."
//
// user.id
func RemoveFirstElementBySeparator(key string, separator string) string {
	idx := strings.Index(key, separator)
	if idx >= 0 {
		return key[idx+1:]
	}
	return key
}
