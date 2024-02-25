package timer

import (
	"github.com/digidudu/mocki/stub"
	"reflect"
	"time"
)

// Sets golang std function  time.Now to return provided time instead of current time
func StubTimeNow(newNow time.Time) {
	timeNow := reflect.ValueOf(time.Now)
	newTimeNow := reflect.ValueOf(func() time.Time {
		return newNow
	})

	stub.StubFunc(timeNow, newTimeNow)
}

// Removes stub from time.Now()
func UnStubTimeNow() {
	timeNow := reflect.ValueOf(time.Now)
	stub.UnStubFunc(timeNow.Pointer())
}
