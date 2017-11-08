package phantom

import (
	"testing"
)

var p *phantom

func TestNew(t *testing.T) {
	p, err := NewPhantom()
	if err != nil {
		t.Errorf("NewPhantom() is error: %v", err)
	}

	if p == nil {
		t.Error("NewPhantom() is nil")
	}
}

func TestExec(t *testing.T) {
	/*
		args := []string{"", "https://www.yahoo.co.jp/", "capture.png"}

		err := p.Exec(args)
		if err != nil {
			t.Error("Exec is error: %v", err)
		}
	*/
}
