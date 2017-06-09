package retry

import (
	"math/rand"
	"net/http"
	"time"
)

var (
	defaultSleep = 500 * time.Millisecond
)

type function func() error

// Do runs the passed function until the number of retries is reached.
func Do(fn function, retries int, sleep time.Duration) error {
	if sleep == 0 {
		sleep = defaultSleep
	}

	if err := fn(); err != nil {
		retries--
		if retries <= 0 {
			return err
		}

		// preventing thundering herd problem (https://en.wikipedia.org/wiki/Thundering_herd_problem)
		sleep += (time.Duration(rand.Int63n(int64(sleep)))) / 2
		time.Sleep(sleep)

		return Do(fn, retries, 2*sleep)
	}

	return nil
}

type httpfunction func() (*http.Response, error)

// DoHTTP runs the passed function until the number of retries is reached.
// It returns *http.Response and error.
func DoHTTP(fn httpfunction, retries int, sleep time.Duration) (*http.Response, error) {
	var res *http.Response

	err := Do(func() error {
		var err error
		res, err = fn()
		return err
	}, retries, sleep)

	return res, err
}
