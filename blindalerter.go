package tigerserver

import (
	"fmt"
	"io"
	"time"
)

// BlindAlerter interface for any Alert creator
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// BlindAlerterFunc allow you to implement BlindAlerter with a function
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter
func (fn BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	fn(duration, amount, to)
}

// Alerter adheres to the ScheduleAlertAt
func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
