package retrygo

import (
	"errors"
	"testing"
	"time"
)

func TestDoRetry(t *testing.T) {
	attemptsCount := 0

	fn := func() error {
		attemptsCount++
		return nil
	}

	err := Do(fn, 2, time.Second)
	if err != nil {
		t.Errorf("retry.Do returned wrong err value: got %v want %v", err, nil)
	}

	if attemptsCount != 1 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 1)
	}
}

func TestDoRetryWithFail(t *testing.T) {
	failErr := errors.New("fail")
	attemptsCount := 0

	fail := func() error {
		attemptsCount++
		return failErr
	}

	err := Do(fail, 2, time.Second)
	if err == nil {
		t.Errorf("retry.Do returned wrong err value: got %v want %v", err, failErr)
	}

	if attemptsCount != 2 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 2)
	}
}
