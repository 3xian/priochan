package priochan

// Chan is an alternate to go channel.
type Chan struct {
	channel        chan interface{}
	sendCompletion func()
}

// NewChan makes an unbufferd Chan.
// Set sendCompletion to nil if you don't want callback.
func NewChan(sendCompletion func()) *Chan {
	return &Chan{
		channel:        make(chan interface{}),
		sendCompletion: sendCompletion,
	}
}

func NewBufferedChan(size int, sendCompletion func()) *Chan {
	return &Chan{
		channel:        make(chan interface{}, size),
		sendCompletion: sendCompletion,
	}
}

func (c *Chan) Send(msg interface{}) {
	c.channel <- msg
	if c.sendCompletion != nil {
		c.sendCompletion()
	}
}

func (c *Chan) Receive() (msg interface{}) {
	return <-c.channel
}
