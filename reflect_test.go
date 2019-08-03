package goconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	testData := []string{
		"ignored",
		"-key1", "value1",
		"--key2", "value2",
		"---key3", "value3",
		"----key4", "value4",
		"ignored",
		"ignored",
		"-key5", "value5",
		"-boolKey",
		"--anotherKey", "anotherValue",
		"--boolKey2", "--boolKey3", "---boolKey4",
		"--key6", "value6",
		"-finalBool",
	}
	expectedFlags := map[string]interface{}{
		"key1":       "value1",
		"key2":       "value2",
		"key3":       "value3",
		"key4":       "value4",
		"key5":       "value5",
		"key6":       "value6",
		"boolKey":    "true",
		"boolKey2":   "true",
		"boolKey3":   "true",
		"boolKey4":   "true",
		"finalBool":  "true",
		"anotherKey": "anotherValue",
	}
	flags := ParseArgs(testData)
	assert.Equal(t, expectedFlags, flags)
}

func TestGetFieldTags(t *testing.T) {
	thing := struct {
		FieldA string `json:"a"`
		FieldB string `json:"b"`
	}{}
	list := getFieldTags(&thing)
	assert.Equal(t, []string{"a", "b"}, list)
}

func TestGetFieldNames(t *testing.T) {
	thing := struct {
		FieldA string `json:"a"`
		FieldB string `json:"b"`
	}{}
	list := getFieldNames(&thing)
	assert.Equal(t, []string{"FieldA", "FieldB"}, list)
}
