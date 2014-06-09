package priochan

import (
	"testing"
	"time"
)

func TestUnbufferdChan(t *testing.T) {
	c := NewChan()
	c.SetSendCompletion(func() { t.Log(time.Now(), "Send finish") })

	go func() {
		time.Sleep(time.Millisecond * 500)
		c.Send(1)
		t.Log(time.Now(), "Send 1")
	}()

	go func() {
		time.Sleep(time.Millisecond * 500)
		c.Send(2)
		t.Log(time.Now(), "Send 2")
	}()

	go func() {
		time.Sleep(time.Millisecond * 500)
		c.Send(3)
		t.Log(time.Now(), "Send 3")
	}()

	for i := 0; i < 3; i++ {
		msg := c.Receive()
		t.Log(time.Now(), "Receive", msg)
		time.Sleep(time.Millisecond * 1000)
	}
}

func TestBufferdChan(t *testing.T) {
}
