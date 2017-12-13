package retry

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

var (
	errFail = errors.New("fail")
)

func TestRetry(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{
			scenario: "do retry",
			function: testDoRetry,
		},
		{
			scenario: "do retry with fail",
			function: testDoRetryWithFail,
		},
		{
			scenario: "do http retry",
			function: testDoHTTPRetry,
		},
		{
			scenario: "do http retry with fail",
			function: testDoHTTPRetryWithFail,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(t)
		})
	}
}

func testDoRetry(t *testing.T) {
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

func testDoRetryWithFail(t *testing.T) {
	errFail := errors.New("fail")
	attemptsCount := 0
	fail := func() error {
		attemptsCount++
		return errFail
	}

	err := Do(fail, 2, time.Second)
	if err == nil {
		t.Errorf("retry.Do returned wrong err value: got %v want %v", err, errFail)
	}

	if attemptsCount != 2 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 2)
	}
}

func testDoHTTPRetry(t *testing.T) {
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

func testDoHTTPRetryWithFail(t *testing.T) {
	attemptsCount := 0
	fn := func() (*http.Response, error) {
		attemptsCount++
		return &http.Response{}, errFail
	}

	_, err := DoHTTP(fn, 2, time.Second)
	if err == nil {
		t.Errorf("retry.DoHTTP returned wrong err value: got %v want %v", err, nil)
	}

	if attemptsCount != 2 {
		t.Errorf("attemptsCount returned wrong count value: got %v want %v", attemptsCount, 2)
	}
}
