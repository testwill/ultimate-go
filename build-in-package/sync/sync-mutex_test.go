package sync

import (
	"fmt"
	"sync"
	"testing"
)

var (
	state = make(map[int]int, 5)
	mutex = &sync.Mutex{} // 读写互斥， 写的时候不让读
)

func stateMutex() {
	mutex.Lock()
	defer mutex.Unlock()
	go func() {
		state[1] = 1
		state[2] = 2
	}()
}

func Test_Mutex(t *testing.T) {
	stateMutex()
	state[1] = 3
	fmt.Println(state)
}

var x = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) {
	// goroutine 里面用的主线程变量，此时需要lock, 即lock 需要操作的变量x
	// mutex.Lock， 因为用到全局变量， 所以需要lock
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}

func Test_Mutex_2(t *testing.T) {
	mutex := &sync.Mutex{}
	wait := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wait.Add(1)
		go increment(wait, mutex)
	}
	wait.Wait()
	fmt.Println(x)
}
