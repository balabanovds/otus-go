package parallel

import (
	"fmt"
	"sync"
)

func RunParallel(slice []func() error, n int, errorsNum int) {
	var wg sync.WaitGroup

	wg.Add(len(slice))
	ch := make(chan error, n)
	for i, f := range slice {
		go func() {
			defer wg.Done()
			fmt.Printf("Run %v", i)
			ch <- f()
		}()
	}

	wg.Wait()
	close(ch)

	errCntr := 0
	for c := range ch {
		if c != nil {
			errCntr++
		}
		if errCntr >= errorsNum {
			return
		}
	}
}
