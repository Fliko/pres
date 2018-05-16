package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Request struct {
	fn func(*Worker) int // The operation to perform.
	c  chan int          // The channel to return the result.
}

type Worker struct {
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
	index    int          // index in the heap
}

type Pool []*Worker

type Balancer struct {
	pool        Pool
	dispatching Pool
	done        chan *Worker
}

func main() {
	var p Pool // setuup variables for balancer
	var d Pool
	var ch = make(chan *Worker) //Communicates between balancer and worker

	for i := 0; i < 10; i++ { // add worker slice to Balancer
		c := make(chan Request, 1600000)
		proc := Worker{c, 0, i}
		go proc.work(ch)
		p = append(p, &proc)
	}

	bal := Balancer{p, d, ch}

	var wrk = make(chan Request) //Talk between requester and balancer

	go bal.balance(wrk) // balance

	go func() { // Make any number of requests
		for j := 0; j < 1600000; j++ {
			go requester(wrk)
		}
	}()

	time.Sleep(time.Minute)
}

func workFn(w *Worker) int {
	time.Sleep(time.Millisecond)
	return w.index
}

func requester(work chan<- Request) {
	nWorker := 10
	c := make(chan int)
	for {
		// Kill some time (fake load).
		time.Sleep(time.Duration(rand.Intn(nWorker)) * time.Millisecond)
		work <- Request{workFn, c} // send request
		fmt.Println(<-c)           // wait for answer
	}
}

func (w *Worker) work(done chan *Worker) {
	for {
		req := <-w.requests // get Request from balancer
		req.c <- req.fn(w)  // call fn and send result
		done <- w           // we've finished this request
	}
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			b.completed(w) // ...so update its info
		}
	}
}

func (p Pool) Less() *Worker {
	min := p[0]
	ind := 0
	for i := 1; i < len(p); i++ {
		if p[i].pending < min.pending {
			min = p[i]
			ind = i
		}
	}
	tmp := ind + 1
	p = append(p[:ind], p[tmp:]...)
	return min
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := b.pool.Less()
	b.dispatching = append(b.dispatching, w)

	// ...send it the task.
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	for i := 1; i < len(b.dispatching); i++ {
		if b.dispatching[i] == w {
			b.dispatching = append(b.dispatching[:i], b.dispatching[i+1:]...)
		}
	}
	b.pool = append(b.pool, w)
}
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
}
