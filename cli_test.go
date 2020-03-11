package tigerserver

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Cable wins\n")
	store := &StubPlayerStore{}
	cli := &CLI{store, in}
	cli.PlayPoker()

	assertPlayerWin(t, store, "Cable")
}
