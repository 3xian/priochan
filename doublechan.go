package priochan

type DoubleChan struct {
	highPrioChan chan interface{}
	lowPrioChan  chan interface{}
	buffer       interface{}
	isBuffered   bool
}

func NewDoubleChan(highPrioChan chan interface{}, lowPrioChan chan interface{}) *DoubleChan {
	return &DoubleChan{highPrioChan, lowPrioChan, nil, false}
}

func (c *DoubleChan) Select() interface{} {
	if c.isBuffered {
		select {
		case highPrioMsg := <-c.highPrioChan:
			return highPrioMsg
		default:
			c.isBuffered = false
			return c.buffer
		}
	}

	select {
	case highPrioMsg := <-c.highPrioChan:
		return highPrioMsg
	case lowPrioMsg := <-c.lowPrioChan:
		select {
		case highPrioMsg := <-c.highPrioChan:
			c.buffer = lowPrioMsg
			c.isBuffered = true
			return highPrioMsg
		default:
			return lowPrioMsg
		}
	}
}
