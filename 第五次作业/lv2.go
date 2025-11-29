package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Run()
}

type Pool struct {
	taskChan chan Task
	wg       sync.WaitGroup
}

func NewPool(workerNum, taskChanCap int) *Pool {
	p := &Pool{
		taskChan: make(chan Task, taskChanCap),
	}

	p.wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func(workerID int) {
			defer p.wg.Done()
			for task := range p.taskChan {

				task.Run()

			}
		}(i)
	}

	return p
}

func (p *Pool) Submit(task Task) {
	p.taskChan <- task
}

func (p *Pool) Wait() {
	close(p.taskChan)
	p.wg.Wait()
}

type PrintTask struct {
	TaskID int    // 任务ID
	Msg    string // 任务信息
}

func (t *PrintTask) Run() {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("任务%d执行成功：%s\n", t.TaskID, t.Msg)
}

type CalcTask struct {
	TaskID int
	A, B   int
	Result chan int
}

func (t *CalcTask) Run() {
	time.Sleep(800 * time.Millisecond)
	sum := t.A + t.B
	t.Result <- sum // 将结果写入通道
	fmt.Printf("计算任务%d执行成功：%d + %d = %d\n", t.TaskID, t.A, t.B, sum)
}

func main() {
	fmt.Println("=== 协程池演示程序启动 ===")

	workerNum := 3
	taskCap := 10
	pool := NewPool(workerNum, taskCap)
	defer pool.Wait()

	fmt.Println("\n=== 提交打印任务 ===")
	for i := 1; i <= 10; i++ {
		task := &PrintTask{
			TaskID: i,
			Msg:    fmt.Sprintf("这是第%d个打印任务", i),
		}
		pool.Submit(task)
	}

	fmt.Println("\n=== 提交计算任务 ===")
	resultChan := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		task := &CalcTask{
			TaskID: i,
			A:      i * 10,
			B:      i * 20,
			Result: resultChan,
		}
		pool.Submit(task)
	}

	go func() {
		defer close(resultChan)
		total := 0
		for sum := range resultChan {
			total += sum
		}
		fmt.Printf("\n=== 所有计算任务总和：%d ===", total)
	}()
	fmt.Println("\n=== 所有任务已提交，等待执行完成 ===")
}
