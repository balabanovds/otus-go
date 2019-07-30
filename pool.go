package pool

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type fn func() error

type job struct {
	num int
	fn
}

type logg struct {
	sync.Mutex
	data []string
}

func newLog() logg {
	return logg{data: []string{}}
}

func (l *logg) add(s string) {
	l.Lock()
	l.data = append(l.data, time.Now().String()+"   "+s)
	l.Unlock()
}
func (l *logg) addf(format string, a ...interface{}) {
	sf := fmt.Sprintf(format, a...)
	l.add(sf)
}

func (l logg) print() {
	for _, v := range l.data {
		fmt.Println(v)
	}
}

// Exec func creates pool of size poolSize and run jobs until maxErrors reached
func Exec(jobs []fn, poolSize int, maxErrors int) {

	l := newLog()

	var wg sync.WaitGroup
	jobsChan := make(chan job, len(jobs))

	// we create channel to hold amount of errors for orchestration workers
	errCh := make(chan int, 1)
	errCh <- 0

	wg.Add(poolSize)
	// create pool of workers
	for i := 0; i < poolSize; i++ {
		go func(i int) {
			defer wg.Done()
			// we subscribe on channel of jobs
			for j := range jobsChan {
				e := <-errCh
				if e < maxErrors {
					var b strings.Builder
					_, _ = fmt.Fprintf(&b, "worker_%d -> processing job_%d", i, j.num)
					err := j.fn()
					if err != nil {
						// if job return error we increase counter
						e++
						_, _ = fmt.Fprintf(&b, " -> got error %d", e)
					}
					l.add(b.String())
				}
				errCh <- e
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
