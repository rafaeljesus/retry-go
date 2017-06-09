package retry

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

var (
	failErr = errors.New("fail")
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

func TestDoHTTPRetry(t *testing.T) {
	attemptsCount := 0
	fn := func() (*http.Response, error) {
		attemptsCount++
		return &http.Response{}, nil
	}

	_, err := DoHTTP(fn, 2, time.Second)
	if err != nil {
		t.Errorf("retry.DoHTTP returned wrong err value: got %v want %v", err, nil)
	}

	if attemptsCount != 1 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 1)
	}
}

func TestDoHTTPRetryWithFail(t *testing.T) {
	attemptsCount := 0
	fn := func() (*http.Response, error) {
		attemptsCount++
		return &http.Response{}, failErr
	}

	_, err := DoHTTP(fn, 2, time.Second)
	if err == nil {
		t.Errorf("retry.DoHTTP returned wrong err value: got %v want %v", err, nil)
	}

	if attemptsCount != 2 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 2)
	}
}
