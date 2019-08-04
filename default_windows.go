package goconfig

import (
	"fmt"
	"os"
)

// DefaultConfigLocation returns the expected config file location for a specific application.
func DefaultConfigLocation(appName string) string {
	return fmt.Sprintf("%s/%s/config.json", os.Getenv("APPDATA"), appName)
}
