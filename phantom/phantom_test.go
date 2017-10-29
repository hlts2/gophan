package phantom

import (
	"testing"
)

var p *phantom

func TestNew(t *testing.T) {
	p = NewPhantom()
	if p == nil {
		t.Error("NewPhantom() is nil")
		return
	}

	if p.binPath == "" {
		t.Error("phantom binPath is Empty")
		return
	}
}

func TestExec(t *testing.T) {
	args := []string{"", "https://www.yahoo.co.jp/", "capture.png"}

	err := p.Exec(args)
	if err != nil {
		t.Error("Exec is error: %v", err)
	}
}
