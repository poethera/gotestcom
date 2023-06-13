package test

import (
	"testing"

	"github.com/poethera/gotestcom/pkg/test"
	"github.com/stretchr/testify/assert"
)

func Test_MathSum(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(test.MathSum(1, 2), 3, "??")

}
