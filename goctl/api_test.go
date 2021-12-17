package goctl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenAndRun(t *testing.T) {
	_, err := genAPI(resource)
	assert.Nil(t, err)
	_, err = runAPISrv(resource)
	assert.Nil(t, err)
	//err = curlSrv(resource)
	//assert.Nil(t, err)
}
