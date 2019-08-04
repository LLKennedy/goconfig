package goconfig

import "fmt"

// DefaultConfigLocation returns the expected config file location for a specific application.
func DefaultConfigLocation(appName string) string {
	return fmt.Sprintf("~/%s/config.json", appName)
}
