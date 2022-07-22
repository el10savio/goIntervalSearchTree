package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPut ...
func TestPut(t *testing.T) {
	var putTestSuite = []struct {
		name         string
		values       []Interval
		expectedList []int
	}{
		{"no values", []Interval{}, []int{}},
		{"single value", []Interval{Interval{1, 2}}, []int{1, 2}},
		{"two values", []Interval{Interval{1, 2}, Interval{3, 4}}, []int{1, 2, 3, 4}},
	}

	for _, testCase := range putTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			intervalTree := New()
			defer intervalTree.Clear()

			intervalTree = putValues(intervalTree, testCase.values)

			assert.Equal(t, testCase.expectedList, intervalTree.toList())
		})
	}
}

func putValues(intervalTree *IntervalTree, intervalList []Interval) *IntervalTree {
	for _, interval := range intervalList {
		fmt.Println(interval)

		err := intervalTree.Put(interval.left, interval.right)
		if err != nil {
			fmt.Println("err:", err)
		}

		fmt.Println("step:", intervalTree.toList())
	}

	fmt.Println("----")

	return intervalTree
}
