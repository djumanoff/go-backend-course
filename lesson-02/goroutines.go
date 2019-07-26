package main

import (
	"fmt"
	"sync"
	//"time"
)

var a []int = []int{}

func main() {
	var mu sync.Mutex
	var muOnce sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(k int){
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			a = append(a, k)
		}(i)
	}

	muOnce.Do(func() {
		fmt.Println("only once")
	})

	muOnce.Do(func() {
		fmt.Println("already twice")
	})

	wg.Wait()
	fmt.Println(a)
}
