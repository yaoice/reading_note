/*
比较通用的做法：通过time Ticker获取一个time的channel
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		DefaultInterval = 1
	)

	done := make(chan bool)
	ticker := time.NewTicker(time.Duration(DefaultInterval) * time.Second)

	fmt.Println("begin!")
	go func() {
		for {
			select {
			case <- done:
				fmt.Println("stop!")
				ticker.Stop()
				return
			case <- ticker.C:
				fmt.Printf("time: %s, msg: trigger the periodic timer.\n", time.Now())
			}
		}
	}()
	time.Sleep(5*time.Second)
	done <- true
	fmt.Println("end!")
}
