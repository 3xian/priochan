package priochan

// DoubleChan is another simple wrapper for double go channels.
type DoubleChan struct {
	highPrioChan chan interface{}
	lowPrioChan  chan interface{}
}

func NewDoubleChan(highPrioChan, lowPrioChan chan interface{}) *DoubleChan {
	return &DoubleChan{highPrioChan, lowPrioChan}
}

// Select returns a message from a go channel with higher priority.
// Take care if any go channels were closed.
func (c *DoubleChan) Select() interface{} {
	select {
	case highPrioMsg, ok := <-c.highPrioChan:
		if ok {
			return highPrioMsg
		} else {
			return <-c.lowPrioChan
		}
	default:
		select {
		case highPrioMsg := <-c.highPrioChan:
			return highPrioMsg
		case lowPrioMsg := <-c.lowPrioChan:
			return lowPrioMsg
		}
	}
}
