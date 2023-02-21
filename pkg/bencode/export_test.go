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
		"\nType: %s\n"+
		"Value: %s\n"+
		"}", typName[o.Typ], o.Val)
}
