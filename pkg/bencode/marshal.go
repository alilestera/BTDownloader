package bencode

import (
	"errors"
	"io"
	"reflect"
)

func Unmarshal(r io.Reader, s interface{}) error {
	o, err := Parse(r)
	if err != nil {
		return err
	}
	p := reflect.ValueOf(s)
	if p.Kind() != reflect.Ptr {
		return errors.New("dest must be a pointer")
	}

	switch o.Typ {
	case BLIST:
		list, _ := o.List()
		l := reflect.MakeSlice(p.Elem().Type(), len(list), len(list))
		p.Elem().Set(l)
		err = unmarshalList(p, list)
		if err != nil {
			return err
		}
	case BDICT:
		//dict, _ := o.Dict()
		//err = unmarshalDict(p, dict)
		//if err != nil {
		//	return err
		//}
	default:
		return errors.New("src code must be struct or slice")
	}
	return nil
}

func unmarshalList(p reflect.Value, list []*BObject) error {
	if p.Kind() != reflect.Ptr || p.Elem().Type().Kind() != reflect.Slice {
		return errors.New("dest must be pointer to slice")
	}
	v := p.Elem()
	if len(list) == 0 {
		return nil
	}

	switch list[0].Typ {
	case BSTR:
		for i, o := range list {
			val, err := o.Str()
			if err != nil {
				return err
			}
			v.Index(i).SetString(val)
		}
	case BINT:
		for i, o := range list {
			val, err := o.Int()
			if err != nil {
				return err
			}
			v.Index(i).SetInt(int64(val))
		}
	case BLIST:
		for i, o := range list {
			val, err := o.List()
			if err != nil {
				return err
			}
			if v.Type().Elem().Kind() != reflect.Slice {
				return ErrTyp
			}
			lp := reflect.New(v.Type().Elem())
			ls := reflect.MakeSlice(v.Type().Elem(), len(val), len(val))
			lp.Elem().Set(ls)
			err = unmarshalList(lp, val)
			if err != nil {
				return err
			}
			v.Index(i).Set(lp.Elem())
		}
	}
	return nil
}
