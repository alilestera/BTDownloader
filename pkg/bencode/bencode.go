package bencode

import (
	"bufio"
	"errors"
	"io"
)

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

var ErrTyp = errors.New("invalid object type")

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

func (o *BObject) Bencode(w io.Writer) int {
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	wLen := 0
	switch o.Typ {
	case BSTR:
		str, _ := o.Str()
		wLen += EncodeString(bw, str)
	case BINT:
		val, _ := o.Int()
		wLen += EncodeInt(bw, val)
	case BLIST:
		bw.WriteByte('l')
		list, _ := o.List()
		for _, elem := range list {
			wLen += elem.Bencode(bw)
		}
		bw.WriteByte('e')
		wLen += 2
	case BDICT:
		bw.WriteByte('d')
		dict, _ := o.Dict()
		for k, v := range dict {
			wLen += EncodeString(bw, k)
			wLen += v.Bencode(bw)
		}
		bw.WriteByte('e')
		wLen += 2
	}
	bw.Flush()
	return wLen
}

func Parse(r io.Reader) (*BObject, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}
	// recursive descent parsing
	bs, err := br.Peek(1)
	if err != nil {
		return nil, err
	}
	var ret BObject
	b := bs[0]
	switch {
	case b >= '0' && b <= '9':
		// parse string
		val, err := DecodeString(br)
		if err != nil {
			return nil, err
		}
		ret.Typ = BSTR
		ret.Val = val
	case b == 'i':
		// parse int
		val, err := DecodeInt(br)
		if err != nil {
			return nil, err
		}
		ret.Typ = BINT
		ret.Val = val
	case b == 'l':
		// parse list
		br.ReadByte()
		var list []*BObject
		for {
			if p, _ := br.Peek(1); p[0] == 'e' {
				br.ReadByte()
				break
			}
			elem, err := Parse(br)
			if err != nil {
				return nil, err
			}
			list = append(list, elem)
		}
		ret.Typ = BLIST
		ret.Val = list
	case b == 'd':
		// parse map
		br.ReadByte()
		dict := map[string]*BObject{}
		for {
			if p, _ := br.Peek(1); p[0] == 'e' {
				br.ReadByte()
				break
			}
			key, err := DecodeString(br)
			if err != nil {
				return nil, err
			}
			val, err := Parse(br)
			if err != nil {
				return nil, err
			}
			dict[key] = val
		}
		ret.Typ = BDICT
		ret.Val = dict
	}
	return &ret, nil
}
