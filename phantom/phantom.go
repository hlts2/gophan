package phantom

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	res "github.com/hlts2/gophan/resource"
)

type phantom struct {
	cmd *exec.Cmd
}

func NewPhantom() (*phantom, error) {
	if err := createBin(); err != nil {
		return nil, err
	}

	p := &phantom{
		cmd: exec.Command(getBinPath()),
	}

	return p, nil
}

/**
  @params args Arguments to pass to phantomjs([0]javascript [1] URL, [2] Output File)
*/
func (p *phantom) Exec(args []string) error {
	p.cmd.Args = args
	if err := p.cmd.Run(); err != nil {
		return err
	}
	return nil
}

func getCachePath() string {
	return filepath.Join(os.TempDir(), "gophan_chache")
}

func getBinPath() string {
	dir := getCachePath()

	var ph string
	if runtime.GOOS == "windows" {
		ph = filepath.Join(dir, "phantomjs.exe")
	} else {
		ph = filepath.Join(dir, "phantomjs")
	}

	return ph
}

func createBin() error {
	binP := getBinPath()

	if _, err := os.Stat(binP); os.IsNotExist(err) {
		name := res.AssetName()

		b, err := res.Asset(name)
		if err != nil {
			return err
		}

		dir := getCachePath()
		os.Mkdir(dir, os.ModePerm)

		ioutil.WriteFile(binP, b, 0600)
	}

	return nil
}
