package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerateToke("sercret", "123")
	spew.Dump(tk)
	panic("stop")
}
