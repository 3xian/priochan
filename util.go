package priochan

import (
	"runtime/debug"
	"testing"
)

func AssertForTest(t *testing.T, actual, expect interface{}) {
	if actual != expect {
		t.Fatalf("\nactual:%v\nexpect:%v\n%s", actual, expect, debug.Stack())
	}
}
