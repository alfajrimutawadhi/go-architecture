package common

import "time"

func CurrentTime() time.Time {
	timeLoc, err := time.LoadLocation("Asia/Jakarta")	// change with your location
	if err != nil {
		return time.Now()
	}
	return time.Now().In(timeLoc)
}