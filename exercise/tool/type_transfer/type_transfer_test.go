package type_transfer

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestTemp(t *testing.T) {
	file := "idl/source.thrift"
	oldPkg, newPkg := "promotion_decision", "promotion_trade_decision"
	methods := []string{"GetPromotions"}

	content := ""
	for _, method := range methods {
		buf, err := GenRequestAndResponseTransferFunc(method, file, oldPkg, newPkg)
		assert.Nil(t, err)
		content += buf
	}

	err := ioutil.WriteFile("out.txt", []byte(content), 0644)
	assert.Nil(t, err)
}
