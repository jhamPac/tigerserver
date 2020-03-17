package tigerserver

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	t.Run("record Cable win from user input", func(t *testing.T) {
		in := strings.NewReader("Cable wins\n")
		store := &StubPlayerStore{}
		cli := NewCLI(store, in, &SpyBlindAlerter{})
		cli.PlayPoker()

		assertPlayerWin(t, store, "Cable")
	})

	t.Run("record Bishop win from user input", func(t *testing.T) {
		in := strings.NewReader("Bishop wins\n")
		store := &StubPlayerStore{}

		cli := NewCLI(store, in, &SpyBlindAlerter{})
		cli.PlayPoker()

		assertPlayerWin(t, store, "Bishop")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Sabertooth wins\n")
		store := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(store, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) <= 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Storm wins\n")
		store := &StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := NewCLI(store, in, blindAlerter)
		cli.PlayPoker()

		cases := []struct {
			expectedScheduleTime time.Duration
			expectedAmount       int
		}{
			{0 * time.Second, 100},
			{10 * time.Second, 200},
			{20 * time.Second, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduleTime), func(t *testing.T) {
				if len(blindAlerter.alerts) <= 1 {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				alert := blindAlerter.alerts[i]

				amountGot := alert.amount
				if amountGot != c.expectedAmount {
					t.Errorf("got amount %d but wanted %d", amountGot, c.expectedAmount)
				}

				gotScheduledTime := alert.scheduledAt
				if gotScheduledTime != c.expectedScheduleTime {
					t.Errorf("got scheduled time of %v but wanted %v", gotScheduledTime, c.expectedScheduleTime)
				}
			})
		}
	})
}
