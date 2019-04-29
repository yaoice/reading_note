/*
自定义一种定时器执行任务的job机制，在time Ticker基础上的升级版，大体思路是：
1. 定义一个Periodic接口类型
2. 定义一个refreshData结构体，实现Periodic接口
3. 定义一个DoPeriodic函数，遍历Periodic类型列表

参考链接：https://github.com/helm/monocular
*/
package main

import (
	"fmt"
	"time"
)

type Periodic interface {
	Do() error
	Frequency() time.Duration
	Name() string
	FirstRun() bool
}

type PeriodicCanceller func()

func DoPeriodic(pSlice []Periodic) PeriodicCanceller {
	doneCh := make(chan struct{})
	for _, p := range pSlice {
		go func(p Periodic) {
			if p.FirstRun() {
				err := p.Do()
				if err != nil {
					fmt.Printf("periodic job ran and returned error (%s)\n", err)
				} else {
					fmt.Printf("periodic job %s ran\n", p.Name())
				}
			}
			ticker := time.NewTicker(p.Frequency())
			for {
				select {
				case <-ticker.C:
					err := p.Do()
					if err != nil {
						fmt.Printf("periodic job ran and returned error (%s)\n", err)
					}
				case <-doneCh:
					ticker.Stop()
					return
				}
			}
		}(p)
	}
	return func() {
		close(doneCh)
	}
}

func NewRefreshData(frequency time.Duration, name string, firstRun bool) Periodic {
	return &refreshData{
		frequency: frequency,
		name:      name,
		firstRun:  firstRun,
	}
}

type refreshData struct {
	frequency time.Duration
	name      string
	firstRun  bool
}

func (r *refreshData) Do() error {
	fmt.Printf("time: %s, %s Do xxx\n", time.Now(), r.name)
	return nil
}

func (r *refreshData) Frequency() time.Duration {
	return r.frequency
}

func (r *refreshData) FirstRun() bool {
	return r.firstRun
}

func (r *refreshData) Name() string {
	return r.name
}

func main() {
	var refreshInterval = 3
	freshness := time.Duration(refreshInterval) * time.Second
	periodicRefresh :=  NewRefreshData(freshness, "refresh", false)
	newPeriodicRefresh :=  NewRefreshData(5*time.Second, "test", true)
	toDo := []Periodic{periodicRefresh, newPeriodicRefresh}
	DoPeriodic(toDo)

	select {
	}
}
