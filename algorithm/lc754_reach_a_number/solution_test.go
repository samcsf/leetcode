package lc754

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
		&TestCase{Args{1}, 1},
		&TestCase{Args{4}, 3},
		&TestCase{Args{5}, 5},
	}
	for _, c := range cases {
		res := ReachNumber(c.input[0].(int))
		pass := reflect.DeepEqual(res, c.expect.(int))
		if !pass {
			t.Errorf("\n Case %v fail, Expect %v Got %v\n", c.input, c.expect, res)
		}
	}
}
