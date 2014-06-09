package priochan

import "sync"

// Chan is an alternate to go channel.
type Chan struct {
	channel             chan interface{}
	sendCompletionMutex sync.RWMutex
	sendCompletion      func()
}

// NewChan makes an unbufferd Chan.
func NewChan() *Chan {
	return &Chan{channel: make(chan interface{})}
}

func NewBufferedChan(size int) *Chan {
	return &Chan{channel: make(chan interface{}, size)}
}

func (c *Chan) Send(msg interface{}) {
	c.channel <- msg
	c.sendCompletionMutex.RLock()
	defer c.sendCompletionMutex.RUnlock()
	if c.sendCompletion != nil {
		c.sendCompletion()
	}
}

func (c *Chan) Receive() (msg interface{}) {
	return <-c.channel
}

// SetSendCompletion assigns a callback function which will be called after Send.
func (c *Chan) SetSendCompletion(f func()) {
	c.sendCompletionMutex.Lock()
	defer c.sendCompletionMutex.Unlock()
	c.sendCompletion = f
}
