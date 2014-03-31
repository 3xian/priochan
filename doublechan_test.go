package priochan

import (
	"testing"
	"time"
)

func TestDoubleChan(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	doubleChan := NewDoubleChan(c1, c2)

	go func() {
		time.Sleep(time.Millisecond * 50)
		for i := 0; i < 3; i++ {
			c1 <- "one"
			time.Sleep(time.Millisecond * 50)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			c2 <- "two"
			time.Sleep(time.Millisecond * 50)
		}
	}()

	expects := []string{"two", "one", "one", "one", "two", "two"}

	for _, expect := range expects {
		msg := doubleChan.Select()
		t.Log(msg)
		if msg != expect {
			t.Error("Expect:", expect)
		}
		time.Sleep(time.Millisecond * 200)
	}
}
