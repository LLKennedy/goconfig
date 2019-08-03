package goconfig

import (
	"fmt"
	"testing"

	"github.com/LLKennedy/goconfig/internal/mocks/fs"
	"github.com/stretchr/testify/assert"
)

type testType struct {
	FieldA string `json:"fieldA"`
	FieldB string `json:"fieldB"`
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
	t.Run("env only", func(t *testing.T) {
		mfs := fs.New()
		mfs.On("Open", fmt.Sprintf(defaultConfigLocation, "test")).Return(nil, fmt.Errorf("cannot open file"))
		defaults := testType{
			FieldA: "defaultA",
		}
		envMap := map[string]string{
			"TEST_FIELDA": "notDefaultA",
		}
		envMapper := func(in string) string { val, _ := envMap[in]; return val }
		err := Load(&defaults, "test", nil, mfs, envMapper)
		expected := testType{
			FieldA: "notDefaultA",
		}
		assert.Equal(t, expected, defaults)
		assert.EqualError(t, err, "cannot open file")
	})
	t.Run("env and JSON", func(t *testing.T) {
		mfs := fs.New(fs.NewFile(fmt.Sprintf(defaultConfigLocation, "test"), []byte(`{
			"fieldB": "notDefaultB"
		}`), nil, nil, true))
		defaults := testType{
			FieldA: "defaultA",
			FieldB: "defaultB",
		}
		envMap := map[string]string{
			"TEST_FIELDA": "notDefaultA",
		}
		envMapper := func(in string) string { val, _ := envMap[in]; return val }
		err := Load(&defaults, "test", nil, mfs, envMapper)
		expected := testType{
			FieldA: "notDefaultA",
			FieldB: "notDefaultB",
		}
		assert.Equal(t, expected, defaults)
		assert.NoError(t, err)
	})
}
