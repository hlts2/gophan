package phantom

import (
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

func getBinPath() string {
	cache := filepath.Join(os.TempDir(), "gophan_chache")

	var ph string
	if runtime.GOOS == "windows" {
		ph = filepath.Join(cache, "phantomjs.exe")
	} else {
		ph = filepath.Join(cache, "phantomjs")
	}

	return ph
}

func createBin() error {
	path := getBinPath()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		name := res.AssetName()

		b, err := res.Asset(name)
		if err != nil {
			return err
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}

		_, err = f.Write(b)
		if err != nil {
			return err
		}
	}

	return nil
}
