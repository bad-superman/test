// +build !linux

package logging

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
