package phantom

type phantom struct {
	binPath string
}

func NewPhantom() (*phantom, error) {
	p := &phantom{
		binPath: generatePhantomPath(),
	}
	return p, nil
}

func (p *phantom) Exec(js string) {

}
