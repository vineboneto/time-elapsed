package timeelapsed

import (
	"testing"
	"time"
)

func TestTimeElapsed(t *testing.T) {

	elapsed := &TimeElapsed{}

	elapsed.Start()
	startTime := elapsed.start

	time.Sleep(2 * time.Second)

	elapsed.End("Description 1")
	endTime := elapsed.end

	elapsedTime := endTime.Sub(startTime)
	if elapsedTime.Seconds() < 1.9 || elapsedTime.Seconds() > 2.1 {
		t.Errorf("Execution time outside of expected range. Expected: 2 seconds, Obtained: %v", elapsedTime)
	}

	elapsed.Start()
	startTime = elapsed.start

	time.Sleep(2 * time.Second)

	elapsed.End("Description 2")
	endTime = elapsed.end

	elapsedTime = endTime.Sub(startTime)
	if elapsedTime.Seconds() < 1.9 || elapsedTime.Seconds() > 2.1 {
		t.Errorf("Execution time outside of expected range. Expected: 2 seconds, Obtained: %v", elapsedTime)
	}
}
