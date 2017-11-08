package resource

import (
	"path/filepath"
	"runtime"
)

func AssetName() string {
	var name string
	if runtime.GOOS == "windows" {
		name = filepath.Join("resource", "windows", "phantomjs.exe")
	} else if runtime.GOOS == "darwin" {
		name = filepath.Join("resource", "darwin", "phantomjs")
	} else {
		name = filepath.Join("resource", "linux", "phantomjs")
	}

	return name
}
