package phantom

import (
	"os"
	"path/filepath"
	"runtime"
)

func generatePhantomPath() string {
	dir, _ := os.Getwd()

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
	dir, _ := os.Getwd()
	return filepath.Join(dir, "..", "resource", "js", "load.js")
}
