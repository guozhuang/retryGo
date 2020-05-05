package main

import (
	"fmt"
	"sync"
	"time"
)

type atomic struct {
	value int
	lock  sync.Mutex
}

func (a *atomic) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomic) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomic
	a.increment()

	go func() {
		a.increment()
	}()

	time.Sleep(time.Microsecond)
	//a是int别名时，go run -race main.go会输出相应的冲突提示
	fmt.Println(a.get()) //go run -race main.go 此时结果不存在读写冲突【说明此时没问题：强制了读锁定和写锁定】
	//如果必须的话，考虑读写锁的使用
}
