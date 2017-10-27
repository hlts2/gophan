package phantom

import (
	"testing"
)

func TestNew(t *testing.T) {
	p := NewPhantom()
	if p == nil {
		t.Error("NewPhantom() is nil")
		return
	}

	if p.binPath == "" {
		t.Error("phantom binPath is Empty")
		return
	}
}
