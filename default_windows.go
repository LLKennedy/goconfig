package goconfig

import (
	"fmt"
	"os"
)

func defaultConfigLocation() string {
	return fmt.Sprintf("%s/%%s/config.json", os.Getenv("APPDATA"))
}
