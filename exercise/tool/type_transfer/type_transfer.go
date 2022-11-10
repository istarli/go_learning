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

const (
	BuilderTypeBasic = iota
	BuilderTypeStruct
	BuilderTypeList
	BuilderTypeMap
)

type BuilderNode struct {
	BuilderType int
	St          *parser.Struct
	Typ         *parser.Type
	Content     string
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

	structBuilderList := []*BuilderNode{{St: root}}

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

func genStructTypeTransferFuncCore(node *BuilderNode, tree map[string]*parser.Thrift, structBuiderList []*BuilderNode) []*BuilderNode {
	oldPkg, newPkg := transPkgName()
	oldType, newType := unpackType(node.Typ, oldPkg), unpackType(node.Typ, newPkg)
	title := genTitle(newPkg, node.Typ)

	var buf strings.Builder
	buf.WriteString(withTapAndLF(0, "func transTo%s(in %s) %s {", title, oldType, newType))
	buf.WriteString(withTapAndLF(1, "if in == nil {"))
	buf.WriteString(withTapAndLF(2, "return nil"))
	buf.WriteString(withTapAndLF(1, "}"))

	switch node.BuilderType {
	case BuilderTypeStruct:
		st := node.St
		if st == nil {
			panic("struct is nil")
		}
		buf.WriteString(withTapAndLF(1, "out := %s.New%s()", newPkg, st.Name))
		for _, field := range st.Fields {
			if isBasicType(field.Type) {
				buf.WriteString(withTapAndLF(1, "out.%s = in.%s", field.Name, field.Name))
			} else {
				if field.Type.Name == "list" {
					builderType := getBuilderType(field.Type.ValueType)
					innerType := unpackType(field.Type.ValueType, newPkg)
					innerTitle := genTitle(newPkg, field.Type.ValueType)

					buf.WriteString(withTapAndLF(1, "out.%s = make([]*%s,0)", field.Name, innerType))
					buf.WriteString(withTapAndLF(1, "for _,item := range in.%s {", field.Name))
					buf.WriteString(withTapAndLF(2, "out.%s = append(out.%s,transTo%s(item))", field.Name, field.Name, innerTitle))
					buf.WriteString(withTapAndLF(1, "}"))

					structBuiderList = append(structBuiderList, &BuilderNode{
						BuilderType: builderType,
						Typ:         field.Type.ValueType,
					})
				} else if field.Type.Name == "map" {
					if field.Type.KeyType.Name != "string" {
						panic("仅支持map<string,x>类型") // todo 仅支持map<string,x>类型
					}
					builderType := getBuilderType(field.Type.ValueType)
					innerType := unpackType(field.Type.ValueType, newPkg)
					innerTitle := genTitle(newPkg, field.Type.ValueType)

					buf.WriteString(withTapAndLF(1, "out.%s = make(map[string]%s)", field.Name, innerType))
					buf.WriteString(withTapAndLF(1, "for key,val := range in.%s {", field.Name))
					buf.WriteString(withTapAndLF(2, "out.%s[key] = transTo%s(val)", field.Name, innerTitle))
					buf.WriteString(withTapAndLF(1, "}"))

					structBuiderList = append(structBuiderList, &BuilderNode{
						BuilderType: builderType,
						Typ:         field.Type.ValueType,
					})
				} else {
					stTitle := genTitle(newPkg, field.Type)

					buf.WriteString(withTapAndLF(1, "out.%s = transTo%s(in.%s)", field.Name, stTitle, field.Name))

					structBuiderList = append(structBuiderList, &BuilderNode{
						BuilderType: BuilderTypeStruct,
						Typ:         field.Type,
						St:          findInnerStruct(field.Type.Name, st, tree),
					})
				}
			}
		}
		buf.WriteString(withTapAndLF(1, "return out"))
		buf.WriteString(withTapAndLF(0, "}\n"))
	case BuilderTypeList:
		panic("todo list")
	case BuilderTypeMap:
		panic("todo map")
	}

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

func transPkgName() (string, string) {
	return oldPkgName, newPkgName
}

// map[string][]*pkg.InnerStruct
func unpackType(typ *parser.Type, pkgName string) string {
	switch typ.Name {
	case "list":
		innerFormat := unpackType(typ.ValueType, pkgName)
		return "[]" + innerFormat
	case "map":
		if typ.KeyType.Name != "string" {
			panic("unsupported map type")
		}
		innerFormat := unpackType(typ.ValueType, pkgName)
		return "map[string]" + innerFormat
	default:
		return fmt.Sprintf("*%s.%s", pkgName, typ.Name)
	}
}

func genTitle(pkg string, typ *parser.Type) string {
	return toCamel(pkg) + genTitleCore(typ)
}
func genTitleCore(typ *parser.Type) string {
	switch typ.Name {
	case "list":
		return "List" + genTitleCore(typ.ValueType)
	case "map":
		return "Map" + genTitleCore(typ.ValueType)
	default:
		return typ.Name
	}
}

func getBuilderType(typ *parser.Type) int {
	if isBasicType(typ) {
		return BuilderTypeBasic
	}
	switch typ.Name {
	case "list":
		return BuilderTypeList
	case "map":
		return BuilderTypeMap
	}
	return BuilderTypeStruct
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

func withTapAndLF(tapNum int, format string, a ...any) string {
	prefix := ""
	for i := 0; i < tapNum; i++ {
		prefix += "\t"
	}
	return prefix + fmt.Sprintf(format, a...) + "\n"
}
