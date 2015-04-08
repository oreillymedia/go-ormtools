package ormtools

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
