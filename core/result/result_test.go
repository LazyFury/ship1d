package result_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"lazyfury.github.com/yoomall-server/core/result"
)

func TestResult(t *testing.T) {
	assert := assert.New(t)

	val := result.Ok(1)

	assert.Equal(1, val.Value)
	assert.Nil(val.Error)

	val2 := result.Err(errors.New("err"))
	assert.Equal("err", val2.Error.Error())

	val3 := result.Err(nil)
	assert.Equal(true, val3.IsErr())

	valOut := 1
	var errOut error
	result.Match(val, func(v int) {
		valOut += v
	}, func(err error) {
		errOut = err
	})

	assert.Equal(2, valOut)
	assert.Nil(errOut)
}
