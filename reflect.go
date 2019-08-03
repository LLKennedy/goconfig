package goconfig

import (
	"reflect"
)

// ParseArgs parses a list of strings as flags and values, assuming all remain as strings.
// It is assumed that you've already cleaned any OS wrapping of strings
func ParseArgs(args []string) map[string]interface{} {
	flags := map[string]interface{}{}
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if len(arg) < 2 || arg[0] != byte('-') {
			// Can't be a flag
			continue
		}
		for len(arg) >= 2 && arg[0] == byte('-') {
			arg = arg[1:]
		}
		if i+1 >= len(args) {
			// Must be a boolean at the end of the args
			flags[arg] = "true"
			continue
		}
		// Next argument exists and this arg is a flag stripped of dashes
		nextArg := args[i+1]
		if len(nextArg) >= 3 && nextArg[:2] == "--" {
			// Next arg is a flag, this is a boolean
			flags[arg] = "true"
			continue
		}
		// No special cases, this is just "-flag value"
		flags[arg] = nextArg
		i++
		continue
	}
	return flags
}

// getFieldTags returns the JSON tags of each field in the Options struct.
// This allows auto-filtering and auto-mapping of runtime flags to config options
func getFieldTags(in interface{}) []string {
	var tags []string
	t := reflect.TypeOf(in)
	for i := 0; i < t.NumField(); i++ {
		tags = append(tags, t.Field(i).Tag.Get("json"))
	}
	return tags
}
