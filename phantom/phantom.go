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

/**
  @params args Arguments to pass to phantomjs([0]javascript [1] URL)
*/
func (p *phantom) Exec(args []string) error {
	if args[0] == "" {
		args[0] = generateJSPath()
	}

	err := exec.Command(p.binPath, args...).Run()
	if err != nil {
		return err
	}

	return nil
}
