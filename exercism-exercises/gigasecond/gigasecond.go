package gigasecond

import (
	"fmt"
	"time"
)

// Adds a gigaseconds to a provided time.
func AddGigasecond(t time.Time) time.Time {
	gigaseconds := 1_000_000_000
	days := gigaseconds / 60 / 60 / 24 // Get days from gigasecond.
	gigaseconds -= days * 24 * 60 * 60 // Remove spent seconds from gigasecond.
	return time.Time.Add(t.AddDate(0, 0, days), time.Second*time.Duration(gigaseconds))
}

func Run() {
	testTime, _ := time.Parse("2006-01-02T15:04:05", "2015-01-24T22:00:00")
	fmt.Println(
		AddGigasecond(testTime),
		"| want: 2046-10-02T23:46:40",
	)
}
