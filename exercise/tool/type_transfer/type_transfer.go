package type_transfer

import (
	"fmt"
	"github.com/samuel/go-thrift/parser"
	"strings"
)

func ParseIDL(filename string) (map[string]*parser.Thrift, string, error) {
	p := &parser.Parser{}
	return p.ParseFile(filename)
}

func GenStructTypeTransferFunc(in *parser.Struct) string {
	funcList :=
}

func genStructTypeTransferFuncCore(in *parser.Struct, funcList []strings.Builder) []strings.Builder {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("func trans%s(in *%s) *%s {", in.Name, in.Name, transPackageName(in.Name)))
	buf.WriteString(fmt.Sprintf("out := %s.New%s()", transPackageName(in.Name), in.Name))
	for _, field := range in.Fields {
		if isBasicType(field.Type) {
			buf.WriteString(fmt.Sprintf("out.%s = in.%s\n", field.Name, field.Name))
		} else {
			buf.WriteString(fmt.Sprintf("out.%s = transfer%S(in)", field.Name, field.Name))
		}
	}
	buf.WriteString("return out\n}\n")
	fmt.Println(buf.String())

	funcList = append(funcList,buf)
	return funcList
}

// isBasicType i32 bool string list<x> map<x,x>
func isBasicType(t *parser.Type) bool {
	switch t.Name {
	case "i32", "bool", "string":
		return true
	case "list":
		return isBasicType(t.ValueType)
	case "map":
		return isBasicType(t.KeyType) && isBasicType(t.ValueType)
	default:
		return false
	}
}

func transPackageName(in string) string {
	return in + "New"
}
