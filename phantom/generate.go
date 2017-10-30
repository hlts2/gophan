package phantom

import (
	"path/filepath"
	"runtime"
)

func generatePhantomPath() string {
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)

	if runtime.GOOS == "darwin" {
		return filepath.Join(dir, "..", "resource", "bin", "darwin", "phantomjs")
	} else if runtime.GOOS == "windows" {
		return filepath.Join(dir, "..", "resource", "bin", "windows", "phantomjs.exe")
	} else {
		if runtime.GOARCH == "386" {
			return filepath.Join(dir, "..", "resource", "bin", "linux386", "phantomjs")
		} else {
			return filepath.Join(dir, "..", "resource", "bin", "linux64", "phantomjs")
		}
	}
}

func generateJSPath() string {
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)
	return filepath.Join(dir, "..", "resource", "js", "load.js")
}
