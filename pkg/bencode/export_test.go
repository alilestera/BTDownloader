package bencode

import "fmt"

var typName = map[BType]string{
	BSTR:  "string",
	BINT:  "int",
	BLIST: "list",
	BDICT: "dict",
}

func (o *BObject) String() string {
	return fmt.Sprintf(""+
		"BObject{"+
		"\nType: %v\n"+
		"Value: %v\n"+
		"}", typName[o.Typ], o.Val)
}
