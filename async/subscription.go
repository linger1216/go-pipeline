package async

type Subscription struct {
	err  error
	done chan struct{}
}

func newSubscription() *Subscription {
	ret := &Subscription{}
	ret.done = make(chan struct{})
	return ret
}

func (s *Subscription) Error() error {
	return s.err
}

func (s *Subscription) _closed(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func (s *Subscription) Unsubscribe() {
	if !s._closed(s.done) {
		s.done <- struct{}{}
		close(s.done)
	}
}

func (s *Subscription) Wait() {
	if !s._closed(s.done) {
		<-s.done
	}
}
