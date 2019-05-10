/*
moby项目提供pubsub简单实现
*/
package main

import (
	"fmt"
	"github.com/moby/moby/pkg/pubsub"
	"strings"
    "sync"
    "time"
)

func main() {
    p := pubsub.NewPublisher(100*time.Millisecond, 10)

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})
	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("docker: https://www.docker.com/")
	time.Sleep(1)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
	    defer func() {
	        wg.Done()
        }()
		fmt.Println("golang topic:", <-golang)
	}()
	go func() {
        defer func() {
            wg.Done()
        }()
		fmt.Println("docker topic:", <-docker)
	}()
    wg.Wait()
}
