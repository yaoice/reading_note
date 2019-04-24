/*
在job_worker基础上扩展，增加DispatchNumControl分发控制数，activeAccount worker活跃数控制，
不过会丢部分job

链接：http://blog.cocosdever.com/2018/08/22/goroutine-channel-Job-Worker-mode
*/
package main
import (
	"fmt"
	"runtime"
	"time"
)
// 定义一些全局常量
var (
	MaxWorker             = 10
	MaxDispatchNumControl = 20
)
// Payload 任务里面的负载
type Payload struct {
	Num int
}
// Job 任务结构体
type Job struct {
	Payload Payload
}
// JobQueue 定义全局Job队列, 新增加的任务就丢进该任务队列即可
var JobQueue chan Job
// WorkerList 工作单元切片
var WorkerList []*Worker
//用于控制并发处理的协程数
var DispatchNumControl chan bool
func Limit(job Job) bool {
	select {
	case <-time.After(time.Millisecond * 100):
		fmt.Println("我很忙")
		return false
	case DispatchNumControl <- true:
		// 任务放入全局任务队列channal
		JobQueue <- job
		return true
	}
}
// Worker 工作者单元, 用于执行Job的单元, 数量有限, 由调度中心分配
type Worker struct {
	WorkerPool chan chan Job //存放JobChan的池子
	JobChan    chan Job
	quit       chan bool
	No         int
}
// NewWorker 创建工作单元
func NewWorker(workerPool chan chan Job, no int) *Worker {
	fmt.Println("创建了工作者", no)
	return &Worker{
		WorkerPool: workerPool,
		JobChan:    make(chan Job),
		quit:       make(chan bool),
		No:         no,
	}
}
// Start 开始工作
func (w *Worker) Start() {
	go func() {
		for {
			// 注册JobChan到工作池中,  然后开始工作循环
			w.WorkerPool <- w.JobChan
			fmt.Println("w.WorkerPool <- w.JobChan | w:", w)
			//如果有工作进来就执行工作, 收到退出信号就退出
			select {
			case job := <-w.JobChan:
				//收到job, 开始工作
				fmt.Println("job := <-w.JobChan")
				fmt.Println(job)
				//完成之后释放控制中心额度
				<-DispatchNumControl
				time.Sleep(5 * time.Second)
			case <-w.quit:
				fmt.Println("<-w.Quit | w:", w)
				return
			}
		}
	}()
}
// Stop 暂停工作
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
// Dispatcher 调度中心, 用于创建工作单元Worker, 安排Worker执行Job
type Dispatcher struct {
	WorkerPool  chan chan Job
	MaxWorkers  int
	ActiveCount int
}
// NewDispatcher 创建调度中心
func NewDispatcher(max int) *Dispatcher {
	return &Dispatcher{
		WorkerPool: make(chan chan Job, max),
		MaxWorkers: max,
	}
}
// Run 根据MaxWorkers, 创建工作者, 同时让工作者运行起来
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
		// 将工作单元存进切片中
		WorkerList[i] = worker
		d.ActiveCount++
	}
	go d.dispatcher()
}
// dispatcher 读取全局job队列, 开始分配任务
func (d *Dispatcher) dispatcher() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				// 从池中找到一个空闲的JobChan, 如果没有空闲的就会堵塞
				jobChan := <-d.WorkerPool
				fmt.Println("jobChan := <-d.WorkerPool")
				//把job丢给工作者
				jobChan <- job
				//每次丢进一个job给工作者之后, 就删除一个工作者, 直到工作者数量维持在5个
				fmt.Println("d.ActiveCount: ", d.ActiveCount)
				if d.ActiveCount > 5 {
					worker := WorkerList[d.ActiveCount-1]
					fmt.Println("worker := WorkerList[d.ActiveCount-1] | worker: ", worker)
					worker.Stop()
					d.ActiveCount--
				}
			}(job)
		}
	}
}
// AddQueue 往全局队列中添加job
func AddQueue(n int) {
	for i := 0; i < n; i++ {
		job := Job{Payload{i}}
		fmt.Println("JobQueue <- job", job)
		// 只有在DispatchNumControl缓冲还未满的时候, 才能将job加入到JobQueue中
		// 因为一旦加入到JobQueue之后, 系统立马会将job从队头取出, 分配一个协程去单独处理后续的工作
		// 为了避免协程数量过多, 所以使用Lmit函数做总体控制
		if Limit(job) {
			fmt.Println("任务成功加入全局队列")
		} else {
			fmt.Println("全局队列已满, 暂不处理任务")
			i--
		}
		fmt.Println("当前协程数:", runtime.NumGoroutine())
		time.Sleep(200 * time.Millisecond)
	}
}
func main() {
	DispatchNumControl = make(chan bool, MaxDispatchNumControl)
	JobQueue = make(chan Job)
	WorkerList = make([]*Worker, 10)
	disp := NewDispatcher(MaxWorker)
	disp.Run()
	time.Sleep(1 * time.Second)
	AddQueue(100)
	fmt.Println()
	time.Sleep(1000 * time.Second)
}