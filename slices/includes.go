package slices

import "fmt"

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	log(fmt.Println("%s: not found %v into the list", pkgNames))
	return false
}
