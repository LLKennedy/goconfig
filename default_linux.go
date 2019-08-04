package goconfig

import "fmt"

// DefaultConfigLocation returns the expected config file location for a specific application.
func DefaultConfigLocation(appName string) string {
	return fmt.Printf("~/%s/config.json", appName)
}
