package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureResponceTime(a)

	bDuration := measureResponceTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponceTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
