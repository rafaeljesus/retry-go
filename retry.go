package retrygo

import (
	"math/rand"
	"time"
)

type function func() error

// Do runs the passed function until the number of retries is reached.
func Do(fn function, retries int, sleep time.Duration) error {
	if err := fn(); err != nil {
		if retries--; retries > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2
			time.Sleep(sleep)
			return Do(fn, retries, 2*sleep)
		}
		return err
	}

	return nil
}
