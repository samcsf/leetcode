package lc40

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type Args []interface{}
type TestCase struct {
	input  Args
	expect interface{}
}

func TestSolution(t *testing.T) {
	cases := []*TestCase{
		&TestCase{Args{[]int{1, 1}, 2}, p2d([]int{1, 1})},
		&TestCase{Args{[]int{1, 1, 2, 5}, 8}, p2d([]int{1, 2, 5})},
		&TestCase{Args{[]int{10, 1, 2, 7, 6, 1, 5}, 8}, p2d(
			[]int{1, 7},
			[]int{2, 6},
			[]int{1, 1, 6},
			[]int{1, 2, 5},
		)},
	}
	for _, c := range cases {
		res := CombinationSum(c.input[0].([]int), c.input[1].(int))
		pass := cmpSliceWithoutOrder(res, c.expect.([][]int))
		if !pass {
			fmt.Printf("Case %v fail\n", c.input)
			fmt.Printf("\tExpect %v\n", c.expect)
			fmt.Printf("\tGot    %v\n", res)
			t.Fail()
		}
	}
}

func p2d(arrs ...[]int) [][]int {
	var res [][]int
	for _, arr := range arrs {
		res = append(res, arr)
	}
	return res
}

func cmpSliceWithoutOrder(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool {
		return len(a[i]) < len(a[j])
	})
	sort.Slice(b, func(i, j int) bool {
		return len(b[i]) < len(b[j])
	})
	return reflect.DeepEqual(a, b)
}
