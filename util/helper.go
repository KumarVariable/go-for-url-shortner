package util

import (
	"fmt"
	"time"
)

// Function to return server uptime
func GetServerUptime() time.Duration {
	return time.Since(SERVER_STARTED_AT)
}

// Function to format time duration into human readable format
// into HH::MM:SS format
func FormatDuration(duration time.Duration) string {

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	// represent time in standard "HH:MM:SS" format
	return fmt.Sprintf("%02d h:%02d m:%02d s", hours, minutes, seconds)

}
