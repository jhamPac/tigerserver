package tigerserver

import "testing"

func TestCLI(t *testing.T) {
	store := &StubPlayerStore{}
	cli := &CLI{store}
	cli.PlayPoker()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win but did not get any")
	}
}
