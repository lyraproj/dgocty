// Package dgocty contains the two function necessary to convert a dgo.Value into a cty.Value and
// vice versa.
package dgocty

import (
	"github.com/lyraproj/dgo/dgo"
	"github.com/lyraproj/dgo/vf"
	"github.com/zclconf/go-cty/cty"
)

// FromCty converts a cty.Value into its corresponding dgo.Value
func FromCty(cv cty.Value) dgo.Value {
	if cv.IsNull() {
		return vf.Nil
	}
	var v dgo.Value
	cvt := cv.Type()
	switch {
	case cvt == cty.String:
		v = vf.String(cv.AsString())
	case cvt == cty.Number:
		fv, _ := cv.AsBigFloat().Float64()
		v = vf.Float(fv)
	case cvt == cty.Bool:
		v = vf.Boolean(cv.True())
	case cvt.IsListType(), cvt.IsTupleType():
		vs := make([]dgo.Value, cv.LengthInt())
		i := 0
		cv.ForEachElement(func(k, v cty.Value) bool {
			vs[i] = FromCty(v)
			i++
			return false
		})
		v = vf.Array(vs)
	case cvt.IsMapType(), cvt.IsObjectType():
		m := vf.MutableMap()
		cv.ForEachElement(func(k, v cty.Value) bool {
			m.Put(FromCty(k), FromCty(v))
			return false
		})
		v = m
	case cvt.IsCapsuleType():
		v = vf.Value(cv.EncapsulatedValue())
	}
	return v
}
