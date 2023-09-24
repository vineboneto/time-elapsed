/**
 * Author: Vinicius Gazolla Boneto
 * File: elapsed.go
 */

package lib

import (
	"log"
	"time"
)

type TimeElapsed struct {
	start time.Time
	end   time.Time
}

func (t *TimeElapsed) Start() {

	t.start = time.Now()
}

func (t *TimeElapsed) End(description string) {
	t.end = time.Now()
	elapsed := t.end.Sub(t.start)

	blue := "\033[34m"
	reset := "\033[0m"

	log.Printf("%s%s:%s %vms\n", blue, description, reset, elapsed.Milliseconds())
}
