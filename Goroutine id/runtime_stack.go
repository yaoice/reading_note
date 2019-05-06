/*
利用runtime.Stack可以获取全部Goroutine的栈信息，官方net/http2库中curGoroutineID函数也采用这种方式获取Goroutine id，
https://github.com/golang/net/blob/master/http2/gotrack.go#L51

获取到Goroutine id，可以方便debug，这个id可以作为唯一标识，将Goroutine中调用的函数层级串联起来；
比较典型的例子：在web框架中，在日志中打印这个id，可以很方便对整个请求过程进行跟踪和分析

参考链接：
- https://liudanking.com/performance/golang-%E8%8E%B7%E5%8F%96-goroutine-id-%E5%AE%8C%E5%85%A8%E6%8C%87%E5%8D%97/
- https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-08-goroutine-id.html
 */
package main

import (
    "bytes"
    "fmt"
    "runtime"
    "strconv"
    "sync"
)

func GetGoid() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}

func main() {
    fmt.Println("main", GetGoid())
    var wg sync.WaitGroup
    for i := 0; i < 20; i++ {
        i := i
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println(i, GetGoid())
        }()
    }
    wg.Wait()
}
