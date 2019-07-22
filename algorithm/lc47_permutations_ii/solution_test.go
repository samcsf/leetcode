package lc47

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
		&TestCase{Args{[]int{1, 1, 2}}, p2d([]int{1, 1, 2}, []int{1, 2, 1}, []int{2, 1, 1})},
	}
	for _, c := range cases {
		res := PermuteUnique(c.input[0].([]int))
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
