package tigerserver

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Cable win from user input", func(t *testing.T) {
		in := strings.NewReader("Cable wins\n")
		store := &StubPlayerStore{}
		cli := NewCLI(store, in)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Cable")
	})

	t.Run("record Bishop win from user input", func(t *testing.T) {
		in := strings.NewReader("Bishop wins\n")
		store := &StubPlayerStore{}

		cli := NewCLI(store, in)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Bishop")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Sabertooth wins\n")
		store := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(store, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})
}
