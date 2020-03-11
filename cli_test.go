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

	if len(store.winCalls) < 1 {
		t.Fatal("expected a win call but did not get any")
	}

	got := store.winCalls[0]
	want := "Cable"

	if got != want {
		t.Errorf("didn't record correct winner, got %q, want %q", got, want)
	}
}
