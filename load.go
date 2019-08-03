package goconfig

import (
	"os"

	"golang.org/x/tools/godoc/vfs"
)

var (
	defaultFileSystem = vfs.OS(".")
	defaultEnvGetter  = os.Getenv
)

type loader struct {
	defaults   interface{}
	appName    string
	flags      map[string]interface{}
	fileSystem vfs.FileSystem
	envGetter  func(string) string
}

// Load loads the default config location on the provided file system, returning defaults for any keys not present in the config file.
// Options start as defaults, then load from environment variables, then a config file, then runtime flags, each overriding the previous.
func Load(defaults interface{}, appName string, flags map[string]interface{}, fs vfs.FileSystem, envAccess func(string) string) (opts interface{}, err error) {
	opts = defaults
	l := &loader{
		defaults:   defaults,
		appName:    appName,
		flags:      flags,
		fileSystem: fs,
		envGetter:  envAccess,
	}
	var withEnv interface{}
	withEnv, err = l.applyEnv(defaults)
	if err == nil {
		var withJSON interface{}
		withJSON, err = l.applyJSON(withEnv)
		if err == nil {
			var withFlags interface{}
			withFlags, err = l.applyFlags(withJSON)
			if err == nil {
				opts = withFlags
			}
		}
	}
	return
}

func (l *loader) applyEnv(partial interface{}) (opts interface{}, err error) {
	opts = partial
	return
}

func (l *loader) applyJSON(partial interface{}) (opts interface{}, err error) {
	opts = partial
	return
}

func (l *loader) applyFlags(partial interface{}) (opts interface{}, err error) {
	opts = partial
	return
}
