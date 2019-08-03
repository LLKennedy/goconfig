package goconfig

import (
	"reflect"
)

// GetFieldTags returns the JSON tags of each field in the Options struct.
// This allows auto-filtering and auto-mapping of runtime flags to config options
func getFieldTags(in interface{}) []string {
	var tags []string
	t := reflect.TypeOf(in)
	for i := 0; i < t.NumField(); i++ {
		tags = append(tags, t.Field(i).Tag.Get("json"))
	}
	return tags
}
