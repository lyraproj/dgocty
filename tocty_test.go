package dgocty_test

import (
	"fmt"
	"testing"

	"github.com/lyraproj/dgo/vf"
	"github.com/lyraproj/dgocty"
)

func ExampleToCty_mixed() {
	v := vf.Map("a", 1, "b", vf.Values("c", 3.14, true))
	c := dgocty.ToCty(v, false)
	fmt.Println(c.GoString())
	// Output: cty.ObjectVal(map[string]cty.Value{"a":cty.NumberIntVal(1), "b":cty.TupleVal([]cty.Value{cty.StringVal("c"), cty.NumberFloatVal(3.14), cty.True})})
}

func ExampleToCty_list() {
	v := vf.Strings("a", "b")
	c := dgocty.ToCty(v, true)
	fmt.Println(c.GoString())
	// Output: cty.ListVal([]cty.Value{cty.StringVal("a"), cty.StringVal("b")})
}

func ExampleToCty_forcedTuple() {
	v := vf.Strings("a", "b")
	c := dgocty.ToCty(v, false)
	fmt.Println(c.GoString())
	// Output: cty.TupleVal([]cty.Value{cty.StringVal("a"), cty.StringVal("b")})
}

func ExampleToCty_map() {
	v := vf.Map("a", 1, "b", 2)
	c := dgocty.ToCty(v, true)
	fmt.Println(c.GoString())
	// Output: cty.MapVal(map[string]cty.Value{"a":cty.NumberIntVal(1), "b":cty.NumberIntVal(2)})
}

func ExampleToCty_forcedObject() {
	v := vf.Map("a", 1, "b", 2)
	c := dgocty.ToCty(v, false)
	fmt.Println(c.GoString())
	// Output: cty.ObjectVal(map[string]cty.Value{"a":cty.NumberIntVal(1), "b":cty.NumberIntVal(2)})
}

func ExampleToCty_nil() {
	fmt.Println(dgocty.ToCty(vf.Nil, false).GoString())
	// Output: cty.NullVal(cty.DynamicPseudoType)
}

func TestToCty_capsule(t *testing.T) {
	s := vf.Sensitive(`secret`)
	c := dgocty.ToCty(s, false)
	if !c.Type().IsCapsuleType() {
		t.Fatal(`sensitive not converted to cty.Capsule`)
	}
	if !s.Equals(dgocty.FromCty(c)) {
		t.Fatal(`sensitive not recreated from cty.Capsule`)
	}
}
