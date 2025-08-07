package data_structures

type Promise[T any] struct {
	channel chan T
}

func NewPromise[T any](routine func() T) *Promise[T] {
	ch := make(chan T)
	go func() {
		defer close(ch)
		ch <- routine()
	}()

	return &Promise[T]{
		channel: make(chan T),
	}
}

func (p *Promise[T]) AwaitPromise() T {
	return <-p.channel
}
