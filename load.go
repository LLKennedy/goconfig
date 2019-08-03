package goconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"golang.org/x/tools/godoc/vfs"
)

var (
	defaultFileSystem = vfs.OS(".")
	defaultEnvGetter  = os.Getenv
)

type loader struct {
	opts       interface{}
	appName    string
	flags      map[string]interface{}
	fileSystem vfs.FileSystem
	envGetter  func(string) string
}

// Load loads the default config location on the provided file system, returning defaults for any keys not present in the config file.
// Options start as defaults, then load from environment variables, then a config file, then runtime flags, each overriding the previous.
func Load(defaults interface{}, appName string, flags map[string]interface{}, fs vfs.FileSystem, envAccess func(string) string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// This isn't elegant, but it is helpful
			err = fmt.Errorf("caught panic: %v\n%s", r, debug.Stack())
		}
	}()
	l := &loader{
		opts:       defaults,
		appName:    appName,
		flags:      flags,
		fileSystem: fs,
		envGetter:  envAccess,
	}
	err = l.applyEnv()
	if err == nil {
		err = l.applyJSON()
		if err == nil {
			err = l.applyFlags()
		}
	}
	return
}

func (l *loader) applyEnv() (err error) {
	capsApp := strings.ToUpper(l.appName)
	for _, fieldName := range getFieldNames(l.opts) {
		capsName := strings.ToUpper(fieldName)
		val := l.envGetter(fmt.Sprintf("%s_%s", capsApp, capsName))
		if val != "" {
			setString(l.opts, fieldName, val)
		}
	}
	return
}

func (l *loader) applyJSON() (err error) {
	var file vfs.ReadSeekCloser
	file, err = l.fileSystem.Open(fmt.Sprintf(defaultConfigLocation, l.appName))
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(l.opts)
	}
	return
}

func (l *loader) applyFlags() (err error) {
	return
}
