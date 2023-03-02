package bencode

import (
	"reflect"
	"testing"
)

var objects = []*BObject{
	{Typ: BLIST, Val: []*BObject{
		{Typ: BINT, Val: 1},
	}},
	{Typ: BLIST, Val: []*BObject{
		{Typ: BINT, Val: 1},
		{Typ: BINT, Val: 2},
	}},
	{Typ: BLIST, Val: []*BObject{
		{Typ: BINT, Val: 1},
		{Typ: BINT, Val: 2},
		{Typ: BINT, Val: 3},
	}},
}

func TestUnmarshalSliceOfSlices(t *testing.T) {
	s := &[]*BObject{}
	p := reflect.ValueOf(s)
	l := reflect.MakeSlice(p.Type().Elem(), len(objects), len(objects))
	p.Elem().Set(l)
	err := unmarshalList(p, objects)
	if err != nil {
		t.Error(err)
	}
}
