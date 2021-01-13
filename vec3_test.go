package vec_test

import (
	"testing"

	"github.com/crushed/vec"
)

func TestVec3Basic(t *testing.T) {
	v := vec.NewVec3(100, 76, 320.0123)
	vstr := v.String()

	if vstr == "Vec3{100, 76, 320.0123}" {
		t.Log("Success: ", vstr)
	} else {
		t.Fail()
	}
}
