# MOCKI

Package that provides func stubbing functionallity.

## Installation

```
go get -u github.com/digidudu/mocki
```

## Pre defined Stubs

### Time

To override time.Now func use `StubTimeNow(newTimeNowVal)` and `UnStubTimeNow()` to revert the change. Can be used in test cases that require checking current time.

```go
package main

import (
	"log"
	"time"

	"github.com/digidudu/mocki/timer"
)

func main() {
	newNow, _ := time.Parse(time.DateOnly, "2000-01-01")
	timer.StubTimeNow(newNow)
	log.Println(time.Now())
}

// Output: 2000-01-01 00:00:00 +0000 UTC
```

## Generic Stub functionallity

To stub function import stub package and pass reflections of target and replacement to `StubFunc(target, replacement)`.

```go
package main

import (
	"log"
	"reflect"

	"github.com/digidudu/mocki/stub"
)

func funcA() {
	log.Println("this is function A")
}

func funcB() {
	log.Println("this is function B")
}

func main() {
	a := reflect.ValueOf(funcA)
	b := reflect.ValueOf(funcB)
	stub.StubFunc(a, b)

	funcA()
}

// Output: this is function B
```
