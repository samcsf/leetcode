package lc77

import (
	"reflect"
	"testing"
)

type Args []interface{}
type TestCase struct {
	input  Args
	expect interface{}
}

func TestSolution(t *testing.T) {
	cases := []*TestCase{
		&TestCase{Args{4, 2}, p2d(
			[]int{1, 2},
			[]int{1, 3},
			[]int{1, 4},
			[]int{2, 3},
			[]int{2, 4},
			[]int{3, 4},
		)},
		&TestCase{Args{5, 3}, p2d(
			[]int{1, 2, 3},
			[]int{1, 2, 4},
			[]int{1, 2, 5},
			[]int{1, 3, 4},
			[]int{1, 3, 5},
			[]int{1, 4, 5},
			[]int{2, 3, 4},
			[]int{2, 3, 5},
			[]int{2, 4, 5},
			[]int{3, 4, 5},
		)},
	}
	for _, c := range cases {
		res := Combine(c.input[0].(int), c.input[1].(int))
		pass := reflect.DeepEqual(res, c.expect.([][]int))
		if !pass {
			t.Errorf("\n Case %v fail, Expect %v Got %v\n", c.input, c.expect, res)
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
