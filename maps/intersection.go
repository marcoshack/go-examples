package maps

func StringIntersection(slices ...[]string) ([]string, bool) {
	count := make(map[string]int)
	for _, slice := range slices {
		for _, key := range slice {
			count[key]++
		}
	}
	duplicated := make([]string, 0, len(count))
	for key, n := range count {
		if n > 1 {
			duplicated = append(duplicated, key)
		}
	}
	return duplicated, len(duplicated) > 0
}

func Intersection[T comparable](slices ...[]T) ([]T, bool) {
	count := make(map[T]int)
	for _, slice := range slices {
		for _, key := range slice {
			count[key]++
		}
	}
	duplicated := make([]T, 0, len(count))
	for key, n := range count {
		if n > 1 {
			duplicated = append(duplicated, key)
		}
	}
	return duplicated, len(duplicated) > 0
}

func HasKeyDuplicationWithCountLoop[K comparable, T any](maps ...map[K]T) bool {
	count := make(map[K]int, 10)
	for _, mapItem := range maps {
		for key := range mapItem {
			count[key]++
		}
	}
	for _, c := range count {
		if c > 1 {
			return true
		}
	}
	return false
}

func HasKeyDuplicationWithEarlyExit[K comparable, T any](maps ...map[K]T) bool {
	count := make(map[K]int, 10)
	for _, mapItem := range maps {
		for key := range mapItem {
			count[key]++
			if count[key] > 1 {
				return true
			}
		}
	}
	return false
}
