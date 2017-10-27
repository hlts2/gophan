package phantom

import (
	"testing"
)

func TestNew(t *testing.T) {
	p, err := NewPhantom()
	if err != nil {
		t.Errorf("NewPhantom() is error: %v", err)
		return
	}

	if p == nil {
		t.Error("NewPhantom() is nil")
		return
	}

	if p.binPath == "" {
		t.Error("phantom binPath is Empty")
		return
	}
}
