package priochan

func NewDoubleChan(highPrioChan <-chan interface{}, lowPrioChan <-chan interface{}) <-chan interface{} {
	mixChan := make(chan interface{})

	go func() {
		defer close(mixChan)

		for {
			select {
			case msg, ok := <-highPrioChan:
				if ok {
					mixChan <- msg
				} else {
					joinChan(lowPrioChan, mixChan)
					return
				}
			default:
				select {
				case msg, ok := <-highPrioChan:
					if ok {
						mixChan <- msg
					} else {
						joinChan(lowPrioChan, mixChan)
						return
					}
				case msg, ok := <-lowPrioChan:
					if ok {
						mixChan <- msg
					} else {
						joinChan(highPrioChan, mixChan)
						return
					}
				}
			}
		}
	}()

	return mixChan
}

func joinChan(in <-chan interface{}, out chan<- interface{}) {
	for {
		if msg, ok := <-in; ok {
			out <- msg
		} else {
			return
		}
	}
}
