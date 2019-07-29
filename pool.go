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
	*status
}

type status struct {
	sync.Mutex
	max     int
	current int
}

func (s status) isMax() bool {
	return s.current >= s.max
}

func (s *status) Incr() {
	s.Lock()
	s.current++
	s.Unlock()
}

func (s status) Current() int {
	return s.current
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

	s := status{max: maxErrors}

	wg.Add(poolSize)
	// create pool of workers
	for i := 0; i < poolSize; i++ {
		go func(i int) {
			defer wg.Done()
			// we subscribe on channel of jobs
			for j := range jobsChan {
				var b strings.Builder
				_, _ = fmt.Fprintf(&b, "worker_%d -> called", i)
				if !j.status.isMax() {
					_, _ = fmt.Fprintf(&b, " -> got job_%d", j.num)
					// if job return error we increase counter
					err := j.fn()
					if err != nil {
						j.status.Incr()
						_, _ = fmt.Fprintf(&b, " -> got error %d", j.status.current)
					}
				}
				l.add(b.String())
			}
		}(i)
	}

	// we are sending jobs to pool, and check counter
	// if max errors reached, we stop
	for i, j := range jobs {
		if s.isMax() {
			break
		}

		jobsChan <- job{num: i, fn: j, status: &s}
		l.add("manager sent job " + strconv.Itoa(i) + " to pool")
	}

	close(jobsChan)

	l.add("manager closed channel")
	wg.Wait()
	l.print()
}
