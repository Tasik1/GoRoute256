package main

import (
	"reflect"
	"testing"
)

// Напишите функцию `Concat`, которая получает несколько слайсов
// и склеивает их в один длинный.
// { {1, 2, 3}, {4, 5}, {6, 7} }  => {1, 2, 3, 4, 5, 6, 7}

func Concat(slices [][]int) []int {
	res := make([]int, 0, 32000)
	for _, slice := range slices {
		for _, val := range slice {
			res = append(res, val)
		}
	}
	return res
}

func TestConcat(t *testing.T) {
	test := []struct {
		slices   [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
		{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
	}

	for _, tc := range test {
		assertEqual(t, tc.expected, Concat(tc.slices))
	}
}

func assertEqual(t *testing.T, expected, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Not equal\nExpected: %v\nActual:   %v\n", expected, actual)
	}
}
