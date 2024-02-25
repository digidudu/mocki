package timer

import (
	"testing"
	"time"
)

func TestStubTimeNow(t *testing.T) {
	// given
	stubDate, _ := time.Parse("2006-01-02", "2000-01-01")

	// when
	StubTimeNow(stubDate)

	//then
	now := time.Now()

	if now != stubDate {
		t.Error("expected datetime different from stubbed now")
	}

	UnStubTimeNow()
}

func TestUnStubTimeNow(t *testing.T) {
	// given
	stubDate, _ := time.Parse("2006-01-02", "2000-01-01")

	// when
	StubTimeNow(stubDate)
	UnStubTimeNow()

	//then
	now := time.Now()

	if !now.After(stubDate) {
		t.Errorf("expected datetime  %s should be after stubbed date %s", now.String(), stubDate.String())
	}
}
