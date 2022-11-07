package type_transfer

import (
	"github.com/samuel/go-thrift/parser"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestTemp(t *testing.T) {
	p := &parser.Parser{}
	m, s, err := p.ParseFile("idl/example.thrift")
	assert.Nil(t, err)

	st := m[s].Structs["ExampleRequest"]
	GenTypeTransferCode(st)
}
