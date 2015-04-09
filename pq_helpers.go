package ormtools

import "strings"

// Arrays
// ---------------------------------------

func FromPqStringArray(arr string) []string {
	split := strings.Split(arr[1:len(arr)-1], ", ")
	return CompactStrings(split)
}
