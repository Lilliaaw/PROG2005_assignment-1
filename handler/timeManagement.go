package handler

import "time"

var startTime time.Time //saves the time in which the service started

/*
*	sets the start time of the service
 */
func SetStartTime() {
	startTime = time.Now()
}

/*
*	gets the duration since start time
 */
func GetDuration() time.Duration {
	return time.Since(startTime)
}
