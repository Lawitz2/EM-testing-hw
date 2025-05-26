package hw8

import "sync"

func UnsafeCounter() int {
	var count int
	var wg sync.WaitGroup

	wg.Add(10)
	for range 10 {
		go func() {
			for range 100000 {
				count++
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return count
}

func SafeCounter() int {
	var count int
	var wg sync.WaitGroup
	var lock sync.Mutex

	wg.Add(10)
	for range 10 {
		go func() {
			for i := 0; i < 100000; i++ {
				lock.Lock()
				count++
				lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return count
}
