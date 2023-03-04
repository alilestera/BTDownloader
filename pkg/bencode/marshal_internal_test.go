package bencode

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
	s := &[][]int{}
	p := reflect.ValueOf(s)
	l := reflect.MakeSlice(p.Type().Elem(), len(objects), len(objects))
	p.Elem().Set(l)
	err := unmarshalList(p, objects)
	if err != nil {
		t.Error(err)
	}
	target := &[][]int{
		{1},
		{1, 2},
		{1, 2, 3},
	}
	assert.Equal(t, s, target, "The two int slice should be equal")
}
