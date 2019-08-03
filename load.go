package goconfig

import (
	"os"

	"golang.org/x/tools/godoc/vfs"
)

var envGetter = os.Getenv

// Load loads the default config location on the provided file system, returning defaults for any keys not present in the config file.
// Options start as defaults, then load from environment variables, then a config file, then runtime flags, each overriding the previous.
func Load(empty interface{}, flags map[string]interface{}, fs vfs.FileSystem, envAccess func(string) string) (opts interface{}, err error) {
	return
}
