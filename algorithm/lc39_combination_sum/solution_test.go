package lc39

import (
	"fmt"
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
		&TestCase{Args{[]int{7, 3, 2}, 18}, p2d(
			[]int{2, 2, 2, 2, 2, 2, 2, 2, 2},
			[]int{2, 2, 2, 2, 2, 2, 3, 3},
			[]int{2, 2, 2, 2, 3, 7},
			[]int{2, 2, 2, 3, 3, 3, 3},
			[]int{2, 2, 7, 7},
			[]int{2, 3, 3, 3, 7},
			[]int{3, 3, 3, 3, 3, 3},
		)},
		&TestCase{Args{[]int{2, 3, 6, 7}, 7}, p2d(
			[]int{2, 2, 3},
			[]int{7},
		)},
	}
	for _, c := range cases {
		res := CombinationSum(c.input[0].([]int), c.input[1].(int))
		fmt.Println(res)
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
