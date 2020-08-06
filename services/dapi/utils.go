package dapi

import (
	"strings"
)

func removeSlashForName(name string) string {
	return strings.Split(name, "/")[1]
}
