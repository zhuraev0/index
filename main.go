package main

func Index(list []int, predicate func(int)bool) int {
	for i, eSlice := range list {
		if predicate(eSlice)  {
			return i
		}
	}
	return -1
}

func Any(list []int, predicate func(int) bool) bool {
	for _, eSlice := range list {
		if predicate(eSlice) {
			return false
		}
	}
	return true
}

func All(list []int, predicate func(int) bool) bool {
	for _, eSlice := range list {
		if !predicate(eSlice) {
			return false
		}
	}
	return true
}

func None(list []int, predicate func(int) bool) bool {
	return Index(list, predicate) == -1
}

func Find(list []int, predicate func(int) bool) int {
	return list[Index(list, predicate)]
}

func main() {

}
