package httpdig

import (
	"testing"
	"time"
)

const (
	testIP   = "8.8.8.8"
	testType = "A"
)

func TestRequest(t *testing.T) {
	_, err := Query(testIP, testType)
	if err != nil {
		t.Errorf("request failed: %s", err.Error())
	}
}

func TestRequestWithNormalTimeout(t *testing.T) {
	timeout := 10 * time.Second
	_, err := QueryWithTimeout(testIP, testType, timeout)
	if err != nil {
		t.Errorf("request with timeout failed: %s", err.Error())
	}
}

func TestRequestWithShortTimeout(t *testing.T) {
	timeout := 10 * time.Microsecond
	_, err := QueryWithTimeout(testIP, testType, timeout)
	if err == nil {
		t.Errorf("request with short timeout should have failed: %s", err.Error())
	}
}
