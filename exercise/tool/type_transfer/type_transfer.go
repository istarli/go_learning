package type_transfer

import (
	"fmt"
	"github.com/samuel/go-thrift/parser"
	"strings"
)

type StructBuilder struct {
	St      *parser.Struct
	Content string
}

func ParseIDL(filename string) (map[string]*parser.Thrift, string, error) {
	p := &parser.Parser{}
	return p.ParseFile(filename)
}

func GenStructTypeTransferFunc(root *parser.Struct, stMap map[string]*parser.Thrift) string {
	structBuilderList := []*StructBuilder{{St: root}}

	pos := 0
	for pos < len(structBuilderList) {
		item := structBuilderList[pos]
		if item.Content == "" {
			structBuilderList = genStructTypeTransferFuncCore(item, stMap, structBuilderList)
		}
		pos++
	}

	var buf strings.Builder
	for _, item := range structBuilderList {
		buf.WriteString(item.Content)
	}

	return buf.String()
}

func genStructTypeTransferFuncCore(node *StructBuilder, tree map[string]*parser.Thrift, structBuiderList []*StructBuilder) []*StructBuilder {
	st := node.St
	var buf strings.Builder

	_, newPkg := transPkgName(st, tree)
	buf.WriteString(fmt.Sprintf("func transTo%s%s(in *%s.%s) *%s.%s {\n", toCamel(newPkg), st.Name, newPkg, st.Name, newPkg, st.Name))
	buf.WriteString(fmt.Sprintf("out := %s.New%s()\n", newPkg, st.Name))
	for _, field := range st.Fields {
		if isBasicType(field.Type) {
			buf.WriteString(fmt.Sprintf("out.%s = in.%s\n", field.Name, field.Name))
		} else if field.Type.Name == "list" {
			valSt := findInnerStruct(field.Type.ValueType.Name, st, tree)
			_, newPkg = transPkgName(valSt, tree)
			buf.WriteString(fmt.Sprintf("out.%s = make([]*%s.%s,0,len(in.%s))\n", field.Name, newPkg, valSt.Name, field.Name))
			buf.WriteString(fmt.Sprintf("for _,item := range in.%s {\n", field.Name))
			buf.WriteString(fmt.Sprintf("  out.%s = append(out.%s,transTo%s%s(item))\n", field.Name, field.Name, toCamel(newPkg), valSt.Name))
			buf.WriteString(fmt.Sprintf("}\n"))
			structBuiderList = append(structBuiderList, &StructBuilder{St: valSt})
		} else if field.Type.Name == "map" {
			if field.Type.KeyType.Name != "string" {
				continue // todo 仅支持map<string,x>类型
			}
			valSt := findInnerStruct(field.Type.ValueType.Name, st, tree)
			_, newPkg = transPkgName(valSt, tree)
			buf.WriteString(fmt.Sprintf("out.%s = make(map[string]*%s.%s)\n", field.Name, newPkg, valSt.Name))
			buf.WriteString(fmt.Sprintf("for key,val := range in.%s {\n", field.Name))
			buf.WriteString(fmt.Sprintf("  out.%s[key] = transTo%s%s(val))\n", field.Name, toCamel(newPkg), valSt.Name))
			buf.WriteString(fmt.Sprintf("}\n"))
			structBuiderList = append(structBuiderList, &StructBuilder{St: valSt})
		} else {
			valSt := findInnerStruct(field.Type.Name, st, tree)
			_, newPkg = transPkgName(valSt, tree)
			buf.WriteString(fmt.Sprintf("out.%s = transTo%s%s(in.%s)\n", field.Name, toCamel(newPkg), valSt.Name, field.Name))
			structBuiderList = append(structBuiderList, &StructBuilder{St: valSt})
		}
	}
	buf.WriteString("return out\n}\n\n")

	node.Content = buf.String()
	return structBuiderList
}

// isBasicType i32 bool string list<x> map<x,x>  include.x
func isBasicType(t *parser.Type) bool {
	switch t.Name {
	case "i32", "bool", "string":
		return true
	case "list":
		return isBasicType(t.ValueType)
	case "map":
		return isBasicType(t.KeyType) && isBasicType(t.ValueType)
	default:
		if strings.Contains(t.Name, ".") {
			return true
		}
		return false
	}
}

func transPkgName(st *parser.Struct, thriftTree map[string]*parser.Thrift) (string, string) {
	pkgName := getPkgName(st, thriftTree)
	return pkgName, pkgName + "_new"
}

// todo 不支持嵌套idl的重名结构体
func findInnerStruct(stName string, root *parser.Struct, thriftTree map[string]*parser.Thrift) *parser.Struct {
	for _, thrift := range thriftTree {
		for _, st := range thrift.Structs {
			if st.Name == stName {
				return st
			}
		}
	}
	return nil
}

func getPkgName(st *parser.Struct, thriftTree map[string]*parser.Thrift) string {
	return "test_pkg"
}

func toCamel(s string) string {
	s = strings.Trim(s, "_")
	if s == "" {
		return ""
	}

	var out []byte
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			continue
		}
		if s[i] == '_' {
			i++
			if i < len(s) {
				out = append(out, toUpper(s[i]))
			}
		} else {
			out = append(out, s[i])
		}
	}
	if len(out) > 0 {
		out[0] = toUpper(out[0])
	}
	return string(out)
}

func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}
	return c
}
