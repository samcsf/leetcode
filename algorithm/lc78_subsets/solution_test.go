package lc78

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
			[]int{},
			[]int{3},
			[]int{2},
			[]int{2, 3},
			[]int{1},
			[]int{1, 3},
			[]int{1, 2},
			[]int{1, 2, 3},
		)},
	}
	for _, c := range cases {
		res := Subsets(c.input[0].([]int))
		// Fixit: []int{} will fails the DeepEqual
		pass := len(res[0]) == 0 && len(c.expect.([][]int)[0]) == 0 && reflect.DeepEqual(res[1:], c.expect.([][]int)[1:])
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
