package dgocty

import (
	"reflect"

	"github.com/lyraproj/dgo/dgo"
	"github.com/lyraproj/dgo/vf"
	"github.com/zclconf/go-cty/cty"
)

// ToCty converts a dgo.Value into its corresponding cty.Value. dgo.Types that are not recognized
// will be converted to cty.Capsule values. dgo.Sensitive values will be unwrapped.
//
// If attemptExplicit is true, then a cty.ListVal will be used instead of a cty.TupleVal and a cty.MapVal will be used
// instead of a cty.ObjectVal in cases when all values are of the same cty.Type.
func ToCty(v dgo.Value, attemptExplicit bool) cty.Value {
	var cv cty.Value
	switch v := v.(type) {
	case dgo.String:
		cv = cty.StringVal(v.GoString())
	case dgo.Integer:
		cv = cty.NumberIntVal(v.GoInt())
	case dgo.Float:
		cv = cty.NumberFloatVal(v.GoFloat())
	case dgo.Boolean:
		cv = cty.BoolVal(v.GoBool())
	case dgo.Array:
		cv = toArray(v, attemptExplicit)
	case dgo.Map:
		cv = toMap(v, attemptExplicit)
	case dgo.Nil:
		cv = cty.NullVal(cty.DynamicPseudoType)
	default:
		cv = toCapsule(v)
	}
	return cv
}

func toArray(v dgo.Array, attemptExplicit bool) cty.Value {
	top := v.Len()
	et := cty.DynamicPseudoType
	useList := attemptExplicit
	cvs := make([]cty.Value, top)
	for i := 0; i < top; i++ {
		ev := ToCty(v.Get(i), attemptExplicit)
		cvs[i] = ev
		if useList {
			if et == cty.DynamicPseudoType {
				et = ev.Type()
			} else if !et.Equals(ev.Type()) {
				useList = false
			}
		}
	}
	if useList {
		return cty.ListVal(cvs)
	}
	return cty.TupleVal(cvs)
}

func toMap(v dgo.Map, attemptExplicit bool) cty.Value {
	et := cty.DynamicPseudoType
	useMap := attemptExplicit
	cvs := make(map[string]cty.Value, v.Len())
	v.EachEntry(func(e dgo.MapEntry) {
		ev := ToCty(e.Value(), attemptExplicit)
		cvs[e.Key().String()] = ev
		if useMap {
			if et == cty.DynamicPseudoType {
				et = ev.Type()
			} else if !et.Equals(ev.Type()) {
				useMap = false
			}
		}
	})
	if useMap {
		return cty.MapVal(cvs)
	}
	return cty.ObjectVal(cvs)
}

func toCapsule(v dgo.Value) cty.Value {
	vt := v.Type()
	vtr := vt.ReflectType()
	t := cty.Capsule(vt.String(), vtr)
	vr := reflect.New(vtr)
	vf.ReflectTo(v, vr.Elem())
	return cty.CapsuleVal(t, vr.Interface())
}
