package main

import "testing"

func TestIndex(t *testing.T) {
	index := []struct {
		name string
		list []int
		predicate func(int)bool
		want int
	}{
		{
			"The item under index {3} exists", []int{1,2,3}, func(i int) bool {
			if i==3 {
				return true
			}
			return false
		}, 2},
		{
			"The item under index {3} doesn't exists", []int{1, 2, 3}, func(i int) bool {
			if i == 4 {
				return true
			}
			return false
		}, -1},
	}
	for _, tt := range index {
		if got := Index(tt.list, tt.predicate); got != tt.want {
			t.Errorf("Index() = %v, want %v", got, tt.want)
		}
	}
}

func TestAny(t *testing.T) {
	any := []struct {
		name string
		list []int
		predicate func(int)bool
		want bool
	}{
		{
			"The item under index [i] exists", []int{98, 28, 34, 32, 77, 90}, func(i int) bool {
			if i<5 {
				return true
			}
			return false
		}, true},
		{
			"The item under index [i] doesn't exists", []int{95, 56, 98, 45, 24}, func(i int) bool {
			if i > 5 {
				return true
			}
			return false
		},  true},
	}
	for _, tt := range any {
		if got := Any(tt.list, tt.predicate); got != tt.want {
			t.Errorf("Any: %v, any() = %v, want: %v", tt.name, got, tt.want)
		}
	}
}

func TestAll(t *testing.T) {
	all := []struct {
		name string
		list []int
		predicate func(int)bool
		want bool
	}{
		{"list is true", []int{10, 10, 10}, func(i int) bool {
			if i == 10 {
				return true
			}
			return false
		}, true},
		{"list is false", []int{10, 20, 30}, func(i int) bool {
			if i == 10 {
				return true
			}
			return false
		}, false},
	}
	for _, tt := range all {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.list, tt.predicate); got != tt.want {
				t.Errorf("All: %v, all() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestNone(t *testing.T) {
	none := []struct {
		name string
		list []int
		predicate func(int)bool
		want bool
	}{
		{"list is true", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i >= 6 {
				return true
			}
			return false
		}, true},
	}
	for _, tt := range none {
		t.Run(tt.name, func(t *testing.T) {
			if got := None(tt.list, tt.predicate); got != tt.want {
				t.Errorf("None: %v, none() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	find := []struct {
		name string
		list []int
		predicate func(int)bool
		want int
	}{
		{"find result", []int{1, 2, 3, 4, 5}, func(i int) bool {
			if i == 3 {
				return true
			}
			return false
		}, 3},
	}
	for _, tt := range find {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.list, tt.predicate); got != tt.want {
				t.Errorf("Find: Find() = %v, want %v", got, tt.want)
			}
		})
	}
}