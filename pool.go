package pool

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type fn func() error

type job struct {
	num int
	fn
}

type errData struct {
	sync.Mutex
	counter int
	max     int
}

func (e errData) isMax() bool {
	return e.counter >= e.max
}

func (e *errData) incr() {
	e.Lock()
	e.counter++
	e.Unlock()
}

// Exec func creates pool of size poolSize and run jobs until maxErrors reached
func Exec(jobs []fn, poolSize int, maxErrors int) {

	l := newLog()

	var wg sync.WaitGroup
	jobsChan := make(chan job, len(jobs))

	// we create channel to hold amount of errors for orchestration workers
	ed := errData{max: maxErrors}

	wg.Add(poolSize)
	// create pool of workers
	for i := 0; i < poolSize; i++ {
		go func(i int) {
			defer wg.Done()
			// we subscribe on channel of jobs
			for j := range jobsChan {
				if !ed.isMax() {
					var b strings.Builder
					_, _ = fmt.Fprintf(&b, "worker_%d -> processing job_%d", i, j.num)
					err := j.fn()
					if err != nil {
						// if job return error we increase counter
						ed.incr()
						_, _ = fmt.Fprintf(&b, " -> got error %d", ed.counter)
					}
					l.add(b.String())
				}
			}
		}(i)
	}

	// we are sending jobs to pool
	for i, j := range jobs {
		jobsChan <- job{num: i, fn: j}
		l.add("manager sent job " + strconv.Itoa(i) + " to pool")
	}

	close(jobsChan)

	l.add("manager closed channel")
	wg.Wait()
	l.print()
}
