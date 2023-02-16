package bencode

import "errors"

type BType uint8

const (
	BSTR  BType = 0x01
	BINT  BType = 0x02
	BLIST BType = 0x03
	BDICT BType = 0x04
)

type BValue interface{}

type BObject struct {
	Typ BType
	Val BValue
}

var ErrTyp = errors.New("wrong object type")

func (o *BObject) Str() (string, error) {
	if o.Typ != BSTR {
		return "", ErrTyp
	}
	return o.Val.(string), nil
}

func (o *BObject) Int() (int, error) {
	if o.Typ != BINT {
		return 0, ErrTyp
	}
	return o.Val.(int), nil
}

func (o *BObject) List() ([]*BObject, error) {
	if o.Typ != BLIST {
		return nil, ErrTyp
	}
	return o.Val.([]*BObject), nil
}

func (o *BObject) Dict() (map[string]*BObject, error) {
	if o.Typ != BDICT {
		return nil, ErrTyp
	}
	return o.Val.(map[string]*BObject), nil
}
