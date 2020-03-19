package tigerserver

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyStdin = &bytes.Buffer{}
var dummyStdout = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record Cable win from user input", func(t *testing.T) {
		in := strings.NewReader("Cable wins\n")
		store := &StubPlayerStore{}
		cli := NewCLI(store, in, dummyStdout, dummyBlindAlerter)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Cable")
	})

	t.Run("record Bishop win from user input", func(t *testing.T) {
		in := strings.NewReader("Bishop wins\n")
		store := &StubPlayerStore{}

		cli := NewCLI(store, in, dummyStdout, dummyBlindAlerter)
		cli.PlayPoker()

		assertPlayerWin(t, store, "Bishop")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Sabertooth wins\n")
		store := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(store, in, dummyStdout, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) <= 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Storm wins\n")
		store := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(store, in, dummyStdout, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		store := &StubPlayerStore{}
		stdout := &bytes.Buffer{}
		cli := NewCLI(store, dummyStdin, stdout, dummyBlindAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := PlayerPrompt

		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}
	})
}
