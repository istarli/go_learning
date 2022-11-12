package type_transfer

import (
	"fmt"
	"github.com/samuel/go-thrift/parser"
	"strings"
)

var (
	oldPkgName string
	newPkgName string
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
	Type        *parser.Type
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

	nodeList := []*BuilderNode{{BuilderType: BuilderTypeStruct, St: root, Type: &parser.Type{Name: root.Name}}}

	pos := 0
	for pos < len(nodeList) {
		item := nodeList[pos]
		nodeList = genStructTypeTransferFuncCore(item, thriftTree, nodeList)
		pos++
	}

	var buf strings.Builder
	for _, item := range nodeList {
		buf.WriteString(item.Content)
	}

	return buf.String()
}

func genStructTypeTransferFuncCore(node *BuilderNode, tree map[string]*parser.Thrift, nodeList []*BuilderNode) []*BuilderNode {
	oldPkg, newPkg := transPkgName()
	oldType, newType := unpackType(node.Type, oldPkg), unpackType(node.Type, newPkg)
	title := genTitle(newPkg, node.Type)

	var buf strings.Builder
	buf.WriteString(withTapAndLF(0, "func transTo%s(in %s) %s {", title, oldType, newType))
	buf.WriteString(withTapAndLF(1, "if in == nil {"))
	buf.WriteString(withTapAndLF(2, "return nil"))
	buf.WriteString(withTapAndLF(1, "}"))

	switch node.BuilderType {
	case BuilderTypeStruct:
		st := node.St
		buf.WriteString(withTapAndLF(1, "out := %s.New%s()", newPkg, st.Name))
		for _, field := range st.Fields {
			if isBasicType(field.Type) {
				buf.WriteString(withTapAndLF(1, "out.%s = in.%s", field.Name, field.Name))
			} else {
				if field.Type.Name == "list" {
					builderType := getBuilderType(field.Type.ValueType)
					innerType := unpackType(field.Type.ValueType, newPkg)
					innerTitle := genTitle(newPkg, field.Type.ValueType)

					buf.WriteString(withTapAndLF(1, "out.%s = make([]%s, 0, len(in.%s))", field.Name, innerType, field.Name))
					buf.WriteString(withTapAndLF(1, "for _, val := range in.%s {", field.Name))
					buf.WriteString(withTapAndLF(2, "out.%s = append(out.%s, transTo%s(val))", field.Name, field.Name, innerTitle))
					buf.WriteString(withTapAndLF(1, "}"))

					newNode := &BuilderNode{
						BuilderType: builderType,
						Type:        field.Type.ValueType,
					}
					if builderType == BuilderTypeStruct {
						newNode.St = findInnerStruct(field.Type.ValueType.Name, tree)
					}
					nodeList = append(nodeList, newNode)

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

					newNode := &BuilderNode{
						BuilderType: builderType,
						Type:        field.Type.ValueType,
					}
					if builderType == BuilderTypeStruct {
						newNode.St = findInnerStruct(field.Type.ValueType.Name, tree)
					}
					nodeList = append(nodeList, newNode)
				} else {
					stTitle := genTitle(newPkg, field.Type)

					buf.WriteString(withTapAndLF(1, "out.%s = transTo%s(in.%s)", field.Name, stTitle, field.Name))

					nodeList = append(nodeList, &BuilderNode{
						BuilderType: BuilderTypeStruct,
						Type:        field.Type,
						St:          findInnerStruct(field.Type.Name, tree),
					})
				}
			}
		}
	case BuilderTypeList:
		builderType := getBuilderType(node.Type.ValueType)
		innerType := unpackType(node.Type.ValueType, newPkg)
		innerTitle := genTitle(newPkg, node.Type.ValueType)

		buf.WriteString(withTapAndLF(1, "out := make([]%s, 0, len(in))", innerType))
		buf.WriteString(withTapAndLF(1, "for _, val := range in {"))
		buf.WriteString(withTapAndLF(2, "out = append(out, transTo%s(val))", innerTitle))
		buf.WriteString(withTapAndLF(1, "}"))

		newNode := &BuilderNode{
			BuilderType: builderType,
			Type:        node.Type.ValueType,
		}
		if builderType == BuilderTypeStruct {
			newNode.St = findInnerStruct(node.Type.ValueType.Name, tree)
		}
		nodeList = append(nodeList, newNode)
	case BuilderTypeMap:
		if node.Type.KeyType.Name != "string" {
			panic("仅支持map<string,x>类型") // todo 仅支持map<string,x>类型
		}
		builderType := getBuilderType(node.Type.ValueType)
		innerType := unpackType(node.Type.ValueType, newPkg)
		innerTitle := genTitle(newPkg, node.Type.ValueType)

		buf.WriteString(withTapAndLF(1, "out := make(map[string]%s)", innerType))
		buf.WriteString(withTapAndLF(1, "for key, val := range in {"))
		buf.WriteString(withTapAndLF(2, "out[key] = transTo%s(val)", innerTitle))
		buf.WriteString(withTapAndLF(1, "}"))

		newNode := &BuilderNode{
			BuilderType: builderType,
			Type:        node.Type.ValueType,
		}
		if builderType == BuilderTypeStruct {
			newNode.St = findInnerStruct(node.Type.ValueType.Name, tree)
		}
		nodeList = append(nodeList, newNode)
	}
	buf.WriteString(withTapAndLF(1, "return out"))
	buf.WriteString(withTapAndLF(0, "}\n"))

	node.Content = buf.String()
	return nodeList
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
func findInnerStruct(stName string, thriftTree map[string]*parser.Thrift) *parser.Struct {
	for _, thrift := range thriftTree {
		for _, st := range thrift.Structs {
			if st.Name == stName {
				return st
			}
		}
	}
	return nil
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
