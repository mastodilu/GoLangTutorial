package main

import "testing"

func TestMySum(t *testing.T) {
	expected := 9
	if x := MySum(2, 3, 4); x != expected {
		t.Errorf("Expected %d but got %d", expected, x)
	}
}

func TestMySum2(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{1, 2, 3, 4, 2, 30}, 42},
		test{[]int{1, 2, 3, 4}, 10},
		test{[]int{1, 2, 3, 4, 3}, 13},
		test{[]int{-10, -7}, -17},
	}

	for i, test := range tests {
		if got := MySum(test.data...); got != test.answer {
			t.Errorf("Expected %d, but got %d at %d\n", test.answer, got, i)
		}
	}
}
