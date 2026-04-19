package main

// 19 april 2026
// CH - 13 : Mutexes

// ========================================== L1 : Mutexes in Go ==========================================
import (
	"sync"
	"time"
	"fmt"
)

// type safeCounter struct {
// 	counts map[string]int
// 	mu     *sync.Mutex
// }

// func (sc safeCounter) inc(key string) {
// 	sc.mu.Lock()
// 	defer sc.mu.Unlock()
// 	sc.slowIncrement(key)
// }

// func (sc safeCounter) val(key string) int {
// 	sc.mu.Lock()
// 	defer sc.mu.Unlock()
// 	return sc.slowVal(key)
// }

// // don't touch below this line

// func (sc safeCounter) slowIncrement(key string) {
// 	tempCounter := sc.counts[key]
// 	time.Sleep(time.Microsecond)
// 	tempCounter++
// 	sc.counts[key] = tempCounter
// }

// func (sc safeCounter) slowVal(key string) int {
// 	time.Sleep(time.Microsecond)
// 	return sc.counts[key]
// }

// ========================================== L2 : Why Is It Called a 'mutex'? ==========================================

// Mutual Exclusion

// ========================================== L3 : Mutex Review ==========================================

// 1

// ========================================== L4 : Mutex Review ==========================================

// To safely access shared resources concurrently

func main() {
	m := map[int]int{}
	mu := &sync.Mutex{}
	go writeLoop(m, mu)
	go readLoop(m, mu)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mu *sync.Mutex) {
	for {
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}
}

func readLoop(m map[int]int, mu *sync.Mutex) {
	for {
		mu.Lock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mu.Unlock()
	}
}

// ========================================== L5 : RW Mutex ==========================================


type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}


// ========================================== L6 : Read/Write Mutex Review ==========================================

// 1

// ========================================== L7 : Read/Write Mutex Review ==========================================

// infinite

// ========================================== L8 : Read/Write Mutex Review ==========================================

// No