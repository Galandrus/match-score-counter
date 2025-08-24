package src

import (
	"fmt"
	"time"
)

func formatTime(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	milliseconds := int(d.Milliseconds()) % 1000
	return fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, milliseconds)
}
