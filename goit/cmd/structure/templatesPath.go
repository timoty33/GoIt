package structure

import (
	"os"
	"path/filepath"
)

func GetTemplatesPath() string {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	baseDir := filepath.Dir(execPath)
	return filepath.Join(baseDir, "templates")
}
