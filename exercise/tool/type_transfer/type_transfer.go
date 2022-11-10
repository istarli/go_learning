package type_transfer

import (
	"fmt"
	"github.com/samuel/go-thrift/parser"
	"strings"
)

var (
	oldPkgName = "member_query"
	newPkgName = "member_query_zg"
)

type StructBuilder struct {
	St      *parser.Struct
	Content string
}

func GenRequestAndResponseTransferFunc(methodName string, filePath, oldPkg, newPkg string) (string, error) {
	p := &parser.Parser{}
	thriftTree, absPath, err := p.ParseFile(filePath)
	if err != nil {
		return "", err
	}

	reqSt := thriftTree[absPath].Structs[methodName+"Request"]
	if reqSt == nil {
		reqSt = thriftTree[absPath].Structs[methodName+"Req"]
	}
	respSt := thriftTree[absPath].Structs[methodName+"Response"]
	if respSt == nil {
		respSt = thriftTree[absPath].Structs[methodName+"Resp"]
	}
	request := GenStructTypeTransferFunc(reqSt, thriftTree, oldPkg, newPkg)
	response := GenStructTypeTransferFunc(respSt, thriftTree, newPkg, oldPkg)

	return request + response, nil
}

func GenStructTypeTransferFunc(root *parser.Struct, thriftTree map[string]*parser.Thrift, oldPkg, newPkg string) string {
	oldPkgName = oldPkg
	newPkgName = newPkg

	structBuilderList := []*StructBuilder{{St: root}}

	pos := 0
	for pos < len(structBuilderList) {
		item := structBuilderList[pos]
		if item.Content == "" {
			structBuilderList = genStructTypeTransferFuncCore(item, thriftTree, structBuilderList)
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

	oldPkg, newPkg := transPkgName(st, tree)
	buf.WriteString(withTapAndLF(0, "func transTo%s%s(in *%s.%s) *%s.%s {", toCamel(newPkg), st.Name, oldPkg, st.Name, newPkg, st.Name))
	buf.WriteString(withTapAndLF(1, "if in == nil {"))
	buf.WriteString(withTapAndLF(2, "return nil"))
	buf.WriteString(withTapAndLF(1, "}"))
	buf.WriteString(withTapAndLF(1, "out := %s.New%s()", newPkg, st.Name))
	for _, field := range st.Fields {
		if isBasicType(field.Type) {
			buf.WriteString(withTapAndLF(1, "out.%s = in.%s", field.Name, field.Name))
		} else if field.Type.Name == "list" {
			valSt := findInnerStruct(field.Type.ValueType.Name, st, tree)
			_, newPkg = transPkgName(valSt, tree)
			buf.WriteString(withTapAndLF(1, "out.%s = make([]*%s.%s,0,len(in.%s))", field.Name, newPkg, valSt.Name, field.Name))
			buf.WriteString(withTapAndLF(1, "for _,item := range in.%s {", field.Name))
			buf.WriteString(withTapAndLF(2, "out.%s = append(out.%s,transTo%s%s(item))", field.Name, field.Name, toCamel(newPkg), valSt.Name))
			buf.WriteString(withTapAndLF(1, "}"))
			structBuiderList = append(structBuiderList, &StructBuilder{St: valSt})
		} else if field.Type.Name == "map" {
			if field.Type.KeyType.Name != "string" {
				panic("仅支持map<string,x>类型") // todo 仅支持map<string,x>类型

			}
			if field.Type.ValueType.Name == "map" || field.Type.ValueType.Name == "list" {
				panic("不支持map、list嵌套类型") // todo 不支持map、list嵌套类型
			}
			valSt := findInnerStruct(field.Type.ValueType.Name, st, tree)
			_, newPkg = transPkgName(valSt, tree)
			buf.WriteString(withTapAndLF(1, "out.%s = make(map[string]*%s.%s)", field.Name, newPkg, valSt.Name))
			buf.WriteString(withTapAndLF(1, "for key,val := range in.%s {", field.Name))
			buf.WriteString(withTapAndLF(2, "out.%s[key] = transTo%s%s(val)", field.Name, toCamel(newPkg), valSt.Name))
			buf.WriteString(withTapAndLF(1, "}"))
			structBuiderList = append(structBuiderList, &StructBuilder{St: valSt})
		} else {
			valSt := findInnerStruct(field.Type.Name, st, tree)
			_, newPkg = transPkgName(valSt, tree)
			buf.WriteString(withTapAndLF(1, "out.%s = transTo%s%s(in.%s)", field.Name, toCamel(newPkg), valSt.Name, field.Name))
			structBuiderList = append(structBuiderList, &StructBuilder{St: valSt})
		}
	}
	buf.WriteString(withTapAndLF(1, "return out"))
	buf.WriteString(withTapAndLF(0, "}\n"))

	node.Content = buf.String()
	return structBuiderList
}

// isBasicType i32 bool string list<x> map<x,x>  include.x
func isBasicType(t *parser.Type) bool {
	switch t.Name {
	case "i32", "i64", "bool", "string":
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
	return oldPkgName, newPkgName
}

//// map[string][]InnerStruct
//func unpackType(typ *parser.Type) (string, string, string) {
//	switch typ.Name {
//	case "i32", "i64", "bool", "string":
//		return "%s", "", typ.Name
//	case "list":
//		innerFormat, innerAssign, innerType := unpackType(typ.ValueType)
//		return "[]" + innerFormat, "[]{" + innerAssign + "}", innerType
//	case "map":
//		if typ.KeyType.Name != "string" {
//			panic("unsupported map type")
//		}
//		innerFormat, innerAssign, innerType := unpackType(typ.ValueType)
//		return "map[string]" + innerFormat, "map[string]{"++"}",innerType
//	default:
//		return "%s", "", typ.Name
//	}
//}

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

func withTapAndLF(tapNum int, format string, a ...any) string {
	prefix := ""
	for i := 0; i < tapNum; i++ {
		prefix += "\t"
	}
	return prefix + fmt.Sprintf(format, a...) + "\n"
}
