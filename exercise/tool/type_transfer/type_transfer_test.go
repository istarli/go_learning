package type_transfer

import (
	"github.com/samuel/go-thrift/parser"
	"github.com/stretchr/testify/assert"
	"io/ioutil"

	"testing"
)

func TestTemp(t *testing.T) {
	p := &parser.Parser{}
	m, s, err := p.ParseFile("idl/source.thrift")
	assert.Nil(t, err)

	request := GenStructTypeTransferFunc(m[s].Structs["QueryMemberBusinessAuthRequest"], m, "member_query", "member_query_zg")
	response := GenStructTypeTransferFunc(m[s].Structs["QueryMemberBusinessAuthResponse"], m, "member_query_zg", "member_query")

	ioutil.WriteFile("out.txt", []byte(request+response), 0644)
}
