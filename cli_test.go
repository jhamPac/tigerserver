package tigerserver

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Cable win from user input", func(t *testing.T) {
		in := strings.NewReader("Cable wins\n")
		store := &StubPlayerStore{}
		cli := &CLI{store, in}
		cli.PlayPoker()

		assertPlayerWin(t, store, "Cable")
	})

	t.Run("record Bishop win from user input", func(t *testing.T) {
		in := strings.NewReader("Bishop wins\n")
		store := &StubPlayerStore{}

		cli := &CLI{store, in}
		cli.PlayPoker()

		assertPlayerWin(t, store, "Bishop")
	})
}
