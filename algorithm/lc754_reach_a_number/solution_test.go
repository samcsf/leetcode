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
		&TestCase{Args{}, xxx},
	}
	for _, c := range cases {
		res := XXX(c.input[0].(xxx), c.input[1].(xxx))
		pass := reflect.DeepEqual(res, c.expect.(xxx))
		if !pass {
			t.Errorf("\n Case %v fail, Expect %v Got %v\n", c.input, c.expect, res)
		}
	}
}
