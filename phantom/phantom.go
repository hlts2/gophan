package phantom

import (
	"os/exec"
)

type phantom struct {
	binPath string
}

func NewPhantom() *phantom {
	p := &phantom{
		binPath: generatePhantomPath(),
	}
	return p
}

func (p *phantom) Exec(jsFile string, args ...string) error {
	var jsPath string
	if jsFile == "" {
		jsPath = generateJSPath()
	} else {
		jsPath = jsFile
	}

	err := exec.Command(p.binPath, jsPath).Run()
	if err != nil {
		return err
	}

	return nil
}
