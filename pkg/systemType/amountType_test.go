package systemType

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	A Amount
}

func TestNewAmountJson(t *testing.T) {
	var (
		t7 int64 = 0
		t0 int64 = 1
		t1 int64 = 100
		t2 int64 = 10000
		t3 int64 = 1000000
		t4 int64 = 100000000
		t5 int64 = 10000000000
		t6 int64 = 1000000000000
	)
	assert.Equal(t, NewAmountInt64(t7).String(), "0")
	assert.Equal(t, NewAmountInt64(t0).String(), "0.0000000001")
	assert.Equal(t, NewAmountInt64(t1).String(), "0.00000001")
	assert.Equal(t, NewAmountInt64(t2).String(), "0.000001")
	assert.Equal(t, NewAmountInt64(t3).String(), "0.0001")
	assert.Equal(t, NewAmountInt64(t4).String(), "0.01")
	assert.Equal(t, NewAmountInt64(t5).String(), "1")
	assert.Equal(t, NewAmountInt64(t6).String(), "100")
	assert.Equal(t, NewAmountInt64(t6).GetInt64(), t6)

	var (
		a    = test{A: NewAmountInt64(t6)}
		aStr = "{\"A\":100}"
	)
	aJstr, err := json.Marshal(a)
	assert.ErrorIs(t, err, nil)
	assert.Equal(t, string(aJstr), aStr)

	var b test
	err = json.Unmarshal([]byte(aStr), &b)
	assert.ErrorIs(t, err, nil)
	assert.Equal(t, NewAmountInt64(t6), b.A)
	fmt.Println(aJstr)
}
