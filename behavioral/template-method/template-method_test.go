package pattern

import (
	"testing"
	"time"
)

func TestWaitForPass(t *testing.T) {
	var m *mazeEntry

	timeout := 500 * time.Millisecond
	m = newEntry(&swingEntryway{}, 500*time.Millisecond)

	start := time.Now()

	m.waitForPass()

	totalDuration := time.Now().Sub(start)

	if totalDuration > (timeout + 100*time.Millisecond) {
		t.Fatalf("program took %v to complete, expected something closer to %v", totalDuration, timeout)
	}
}
