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

// BlindAlerterFunc type for any func instead of a stuct
type BlindAlerterFunc func(duration time.Duration, amount int)

// ScheduleAlertAt implements the type
func (fn BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	fn(duration, amount)
}

// StdOutAlerter adheres to the ScheduleAlertAt
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
