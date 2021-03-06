package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*func worker(id int, ch chan int) {
	for {
		//从channel中读数据
		fmt.Printf("worker %d data is %c\n", id, <-ch)
	}
}

func main() {

	//channel+worker交互
	var channels [10]chan int

	//进行任务的基本分发
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)

		go worker(i, channels[i])
	}

	//进行数据填充
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //单引号的字符使用
	}

	//重复填充
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i //单引号的字符使用
	}

	time.Sleep(time.Microsecond * 1000)
}*/

//使用生成器的方式来对应相应worker的channel创建
//新增箭头用来明确外部使用chan的用途【规范使用】
/*func createWorker(id int) chan<- int {
	ch := make(chan int)
	go func() {
		for {
			//这种实现模式也是为啥golang中死循环非常多的原因
			fmt.Printf("worker %d data is %c\n", id, <-ch)
		}
	}()
	return ch
}

func main() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //单引号的字符使用
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i //单引号的字符使用
	}

	time.Sleep(time.Microsecond * 1000)
}

//buffer channel

//手动回收goroutine的原因：接收者判断channel接收到是否done，来决定是否进行回收该关联的goroutine
//done模式和background的模式来形成调用结构的统一管理
*/

//更加工程化的实现模式：

//新增支持done的worker的标准结构
/*type worker struct {
	in   chan int
	done chan bool
}

//done模式的使用
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

//通过worker的完成通知来通知上一级goroutine进行完成
func doWork(id int, ch <-chan int, done chan bool) {
	//先进行接收
	for n := range ch {
		//这种实现模式也是为啥golang中死循环非常多的原因
		fmt.Printf("worker %d data is %c\n", id, n)

		done <- true //如果这样触发的话，因为下面重复写入chan，但是第一次的done没有完成接收，因为接收者的逻辑还在写入的逻辑的下面
		//导致下一次写入数据始终被阻塞，也无法消费done的chan，导致deadlock
		//因为是只有十个worker，下次的写入之前就需要消费
		go func() {
			done <- true
		}()
	}
}

func main() {
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		//逐个任务的触发和执行完成才有下一步，相当于所有的任务顺序执行
		//<- workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		//<- workers[i].done
	}

	for _, worker := range workers {
		//因为填写了两次，所以两次done,这样实现就不是完全同步完成了
		<-worker.done
		<-worker.done //实际上这里的阻塞就完成了整体任务的所有等待时间
	}

	//实际上这样固定数量完成任务直接使用waitGroup即可

	//time.Sleep(time.Microsecond * 1000)//将无脑休眠去掉
}*/

//select的使用[for+select]
func generator() chan int {
	out := make(chan int)

	go func() {
		for {
			num := rand.Intn(1000)
			//目的是形成两路速度不同的channel
			time.Sleep(time.Microsecond * time.Duration(num))

			//进行chan写入
			out <- num
		}
	}()

	return out
}

func main() {
	//生成两路chan
	ch1 := generator()
	ch2 := generator()

	//time.Tick进行定时器处理分支【例如作为查看积压队列情况】
	tick := time.Tick(time.Second)
	for {
		//同时这里可以将拿到的数据再发给别的位置
		select {
		//也可以结合done模式和超时模式来进行处理【添加计时器来进行判断】
		case n := <-ch1:
			fmt.Printf("channel1:%d\n", n)

		case m := <-ch2:
			fmt.Printf("channel2:%d\n", m)
		//超时分支处理【进入select的耗时判断】
		case <-time.After(1000 * time.Microsecond):
			fmt.Println("timeout")

		//定时器分支【例如输出消费挤压情况】
		case <-tick:
			fmt.Println("this is")
		}

	}
}
