package ormtools

import (
	"strconv"
	"strings"
)

// Unique
// ---------------------------------------

func UniqueStrings(oldStrings []string) []string {
	var newStrings []string
Outer:
	for _, oldString := range oldStrings {
		for _, newString := range newStrings {
			if oldString == newString {
				continue Outer
			}
		}
		newStrings = append(newStrings, oldString)
	}
	return newStrings
}

func UniqueInts(oldInts []int) []int {
	var newInts []int
Outer:
	for _, oldInt := range oldInts {
		for _, newInt := range newInts {
			if oldInt == newInt {
				continue Outer
			}
		}
		newInts = append(newInts, oldInt)
	}
	return newInts
}

func UniqueInt64s(oldInts []int64) []int64 {
	var newInts []int64
Outer:
	for _, oldInt := range oldInts {
		for _, newInt := range newInts {
			if oldInt == newInt {
				continue Outer
			}
		}
		newInts = append(newInts, oldInt)
	}
	return newInts
}

// Map
// ---------------------------------------

func StringsToInts(oldStrings []string) (newints []int) {
	for _, s := range oldStrings {
		si, err := strconv.Atoi(s)
		if err == nil {
			newints = append(newints, si)
		}
	}
	return newints
}

func IntsToStrings(oldInts []int) (newStrings []string) {
	for _, i := range oldInts {
		newStrings = append(newStrings, strconv.Itoa(i))
	}
	return newStrings
}

// Compact
// ---------------------------------------

func CompactStrings(oldStrings []string) []string {
	var newStrings []string
	for _, oldString := range oldStrings {
		if oldString != "" {
			newStrings = append(newStrings, oldString)
		}
	}
	return newStrings
}

// Trim
// ---------------------------------------

func TrimStrings(oldStrings []string, cutset string) []string {
	for i, oldString := range oldStrings {
		oldStrings[i] = strings.Trim(oldString, cutset)
	}
	return oldStrings
}
