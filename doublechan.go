package priochan

type DoubleChan struct {
	highPrioChan chan interface{}
	lowPrioChan  chan interface{}
}

func NewDoubleChan(highPrioChan chan interface{}, lowPrioChan chan interface{}) *DoubleChan {
	return &DoubleChan{highPrioChan, lowPrioChan}
}

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
