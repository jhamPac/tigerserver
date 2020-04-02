package tigerserver

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		store := &StubPlayerStore{}
		game := NewTexas(blindAlerter, store)

		game.Start(5, ioutil.Discard)

		cases := []ScheduledAlert{
			{at: 0 * time.Second, amount: 100},
			{at: 10 * time.Minute, amount: 200},
			{at: 20 * time.Minute, amount: 300},
			{at: 30 * time.Minute, amount: 400},
			{at: 40 * time.Minute, amount: 500},
			{at: 50 * time.Minute, amount: 600},
			{at: 60 * time.Minute, amount: 800},
			{at: 70 * time.Minute, amount: 1000},
			{at: 80 * time.Minute, amount: 2000},
			{at: 90 * time.Minute, amount: 4000},
			{at: 100 * time.Minute, amount: 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		ba := &SpyBlindAlerter{}
		store := &StubPlayerStore{}
		game := NewTexas(ba, store)

		game.Start(7, ioutil.Discard)

		cases := []ScheduledAlert{
			{at: 0 * time.Second, amount: 100},
			{at: 12 * time.Minute, amount: 200},
			{at: 24 * time.Minute, amount: 300},
			{at: 36 * time.Minute, amount: 400},
		}

		checkSchedulingCases(cases, t, ba)
	})
}

func TestGame_Finish(t *testing.T) {
	ba := &SpyBlindAlerter{}
	store := &StubPlayerStore{}
	game := NewTexas(ba, store)

	winner := "Magneto"

	game.Finish(winner)
	assertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []ScheduledAlert, t *testing.T, blindAlerter *SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(blindAlerter.alerts) <= 1 {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
			}

			got := blindAlerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}
