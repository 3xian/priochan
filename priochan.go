package priochan

type PrioChan struct {
}

/*

func NewPrioChan() *PrioChan {
	// TODO
}

func NewPrioChanWithSlice() *PrioChan {
	// TODO
}

func (pc* PrioChan) Push() interface{} {
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

func (pc* PrioChan) Select() interface{} {
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
*/
