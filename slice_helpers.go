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
