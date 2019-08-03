package goconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testType struct {
	FieldA string `json:"fieldA"`
}

func TestLoad(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		defaults := testType{
			FieldA: "defaultA",
		}
		envMap := map[string]string{
			"TEST_FIELDA": "notDefaultA",
		}
		envMapper := func(in string) string { val, _ := envMap[in]; return val }
		err := Load(defaults, "test", nil, nil, envMapper)
		expected := testType{
			FieldA: "defaultA",
		}
		assert.Equal(t, expected, defaults)
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "caught panic: reflect: Elem of invalid type")
			assert.Contains(t, err.Error(), "reflect.go")
		}
	})
	t.Run("basics", func(t *testing.T) {
		defaults := testType{
			FieldA: "defaultA",
		}
		envMap := map[string]string{
			"TEST_FIELDA": "notDefaultA",
		}
		envMapper := func(in string) string { val, _ := envMap[in]; return val }
		err := Load(&defaults, "test", nil, nil, envMapper)
		expected := testType{
			FieldA: "notDefaultA",
		}
		assert.Equal(t, expected, defaults)
		assert.NoError(t, err)
	})
}
