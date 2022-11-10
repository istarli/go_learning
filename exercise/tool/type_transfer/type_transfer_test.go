package type_transfer

import (
	"github.com/samuel/go-thrift/parser"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestTemp(t *testing.T) {
	file := "idl/source.thrift"
	oldPkg, newPkg := "channel_router", "channel_router_zg"
	methods := []string{"QuerySupportedBankCard"}

	content := ""
	for _, method := range methods {
		buf, err := GenRequestAndResponseTransferFunc(method, file, oldPkg, newPkg)
		assert.Nil(t, err)
		content += buf
	}

	err := ioutil.WriteFile("out.txt", []byte(content), 0644)
	assert.Nil(t, err)
}

func TestTemp2(t *testing.T) {
	file := "idl/source.thrift"
	oldPkg, newPkg := "channel_router", "channel_router_zg"
	structName := "BankCardInfo"

	p := &parser.Parser{}
	thriftTree, absPath, err := p.ParseFile(file)
	assert.Nil(t, err)

	st := thriftTree[absPath].Structs[structName]
	buf := GenStructTypeTransferFunc(st, thriftTree, oldPkg, newPkg)

	err = ioutil.WriteFile("out.txt", []byte(buf), 0644)
	assert.Nil(t, err)
}
