package donemerge

func Or(channels ...<-chan interface{}) <-chan interface{} {
	doneChan := make(chan interface{})

	for _, ch := range channels {
		go func(channel <-chan interface{}) {
			doneChan <- <-channel
		}(ch)

	}

	return doneChan
}
