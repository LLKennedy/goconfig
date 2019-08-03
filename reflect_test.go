package goconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFieldTags(t *testing.T) {
	thing := struct {
		FieldA string `json:"a"`
		FieldB string `json:"b"`
	}{}
	list := getFieldTags(thing)
	assert.Equal(t, []string{"a", "b"}, list)
}
