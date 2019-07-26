package pool

import (
	"log"
	"sync"
)

type mutexCounter struct {
	sync.Mutex
	counter int
}

type fn func() error

type job struct {
	num int
	fn
}

// Exec func creates pool of size poolSize and run jobs until maxErrors reached
func Exec(jobs []fn, poolSize int, maxErrors int) {
	var wg sync.WaitGroup
	ch := make(chan job)

	// we assign counter to maxErrors to reduce it in pool
	c := mutexCounter{counter: maxErrors}

	wg.Add(poolSize)
	// create pool of workers
	for i := 0; i < poolSize; i++ {
		go func(i int) {
			// we subscribe on channel of jobs
			for j := range ch {
				log.Printf("worker %d, got job %d\n", i, j.num)
				// if job return error we decrease counter
				if j.fn() != nil {
					c.Lock()
					c.counter--
					log.Printf("worker %d, job %d got error, counter is %d\n", i, j.num, c.counter)
					c.Unlock()
				}
			}
			log.Printf("worker %d stopped\n", i)
			wg.Done()
		}(i)
	}

	// we are sending jobs to pool, and check counter
	// if max errors reached, we stop
	var i int
	for i < len(jobs) {
		ch <- job{num: i, fn: jobs[i]}
		log.Printf("manager sent job %d to pool\n", i)
		i++
		c.Lock()
		if c.counter == 0 {
			c.Unlock()
			log.Println("max errors reached")
			break
		}
		c.Unlock()
	}

	close(ch)
	log.Println("manager closed channel")
	wg.Wait()
	log.Printf("jobs done %d of %d\n", i, len(jobs))
}
