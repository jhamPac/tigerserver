package tigerserver

import (
	"fmt"
	"os"
	"time"
)

// BlindAlerter interface for any Alert creator
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// BlindAlerterFunc allow you to implement BlindAlerter with a function
type BlindAlerterFunc func(duration time.Duration, amount int)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter
func (fn BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	fn(duration, amount)
}

// StdOutAlerter adheres to the ScheduleAlertAt
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
