package lc46

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
		&TestCase{Args{[]int{1, 2, 3}}, p2d(
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			[]int{2, 1, 3},
			[]int{2, 3, 1},
			[]int{3, 1, 2},
			[]int{3, 2, 1},
		)},
		&TestCase{Args{[]int{1, 1}}, p2d(
			[]int{1, 1},
			[]int{1, 1},
		)},
	}
	for _, c := range cases {
		res := Permute(c.input[0].([]int))
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
