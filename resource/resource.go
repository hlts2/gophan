package resource

import (
	"path/filepath"
	"runtime"
)

func AssetName() string {
	var name string

	switch runtime.GOOS {
	case "windows":
		name = filepath.Join("resource", "windows", "phantomjs.exe")
	case "darwin":
		name = filepath.Join("resource", "darwin", "phantomjs")
	default:
		name = filepath.Join("resource", "linux", "phantomjs")
	}

	return name
}
