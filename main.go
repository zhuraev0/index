package main

func Index(list []int, predicate func(int)bool) int {
	for i, eSlice := range list {
		if predicate(eSlice)  {
			return i
		}
	}
	return -1
}
func All(list []int, predicate func(int) bool) bool {
	return Index(list, func(i int) bool {
		return !predicate(i)
	}) == -1
}

func Any(list []int, predicate func(int) bool) bool {
	return Index(list, predicate) != 1
}

func None(list []int, predicate func(int) bool) bool {
	return Index(list, predicate) == -1
}

func Find(list []int, predicate func(int) bool) int {
	return list[Index(list, predicate)]
}

func main() {

}
