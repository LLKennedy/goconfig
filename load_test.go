package goconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	loaded, err := Load(nil, "", nil, nil, nil)
	assert.Nil(t, loaded)
	assert.NoError(t, err)
}
