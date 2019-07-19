package lc26

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
		&TestCase{Args{[]int{1, 1, 2}}, []int{1, 2}},
		&TestCase{Args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, []int{0, 1, 2, 3, 4}},
	}
	for _, c := range cases {
		res := RemoveDuplicates(c.input[0].([]int))
		pass := reflect.DeepEqual(c.input[0].([]int)[:res], c.expect)
		if !pass {
			t.Errorf("\n Case %v fail, Expect %v Got %v\n", c.input, c.expect, c.input[0].([]int)[:res])
		}
	}
}
