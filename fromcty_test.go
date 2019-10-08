package dgocty_test

import (
	"fmt"
	"testing"

	"github.com/lyraproj/dgo/dgocty"
	"github.com/lyraproj/dgo/vf"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

func ExampleFromCty_mixed() {
	v, err := gocty.ToCtyValue(map[string]interface{}{"a": 1, "b": []interface{}{"c", 3.14, true}},
		cty.Object(map[string]cty.Type{"a": cty.Number, "b": cty.Tuple([]cty.Type{cty.String, cty.Number, cty.Bool})}))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dgocty.FromCty(v))
	}
	// Output: {"a":1.0,"b":["c",3.14,true]}
}

func TestFromCty_null(t *testing.T) {
	var s []string
	v, err := gocty.ToCtyValue(&s, cty.List(cty.String))
	if err != nil {
		t.Fatal(err)
	}
	if !vf.Nil.Equals(dgocty.FromCty(v)) {
		t.Fatal(`cty value did not produce Nil from zero value`)
	}
}
