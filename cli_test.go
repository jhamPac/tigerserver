package tigerserver

import (
	"bytes"
	"strings"
	"testing"
)

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyStdin = &bytes.Buffer{}
var dummyStdout = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record Cable win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCable wins\n")
		store := &StubPlayerStore{}
		blindAlerter := dummyBlindAlerter
		game := NewGame(blindAlerter, store)

		cli := NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Cable")
	})

	t.Run("record Bishop win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nBishop wins\n")
		store := &StubPlayerStore{}
		blindAlerter := dummyBlindAlerter
		game := NewGame(blindAlerter, store)

		cli := NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Bishop")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Sabertooth wins\n")
		store := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		game := NewGame(blindAlerter, store)

		cli := NewCLI(in, dummyStdout, game)
		cli.PlayPoker()

		if len(blindAlerter.alerts) <= 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})
}
