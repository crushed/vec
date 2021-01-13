package vec_test

import (
	"testing"

	"github.com/crushed/vec"
)

func TestVec2Basic(t *testing.T) {
	v := vec.NewVec2(10, 16)
	vstr := v.String()

	if vstr == "Vec2{10, 16}" {
		t.Log("Success", vstr)
	} else {
		t.Fail()
	}
}
