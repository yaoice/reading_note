/*
etcd client v3版本中的concurrency包实现了分布式锁，
大致思路是：
1. createReVision最小的客户端获得锁
2. createReVision越小越早获得锁，部分关键代码如下

等待比当前客户端创建的key的revision小的key的客户端释放锁
// wait for deletion revisions prior to myKey
hdr, werr := waitDeletes(ctx, client, m.pfx, m.myRev-1)

参考链接: https://github.com/etcd-io/etcd/blob/master/clientv3/example_kv_test.go
*/

package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "time"

    "github.com/coreos/etcd/clientv3"
    "github.com/coreos/etcd/clientv3/concurrency"
)

var (
    endpoints = "http://10.10.10.10:52379"
)


func main() {
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   strings.Split(endpoints, ","),
        DialTimeout: 3 * time.Second,
    })
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    // create two separate sessions for lock competition
    s1, err := concurrency.NewSession(cli)
    if err != nil {
        log.Fatal(err)
    }
    defer s1.Close()
    m1 := concurrency.NewMutex(s1, "/my-lock/")

    s2, err := concurrency.NewSession(cli)
    if err != nil {
        log.Fatal(err)
    }
    defer s2.Close()
    m2 := concurrency.NewMutex(s2, "/my-lock/")

    // acquire lock for s1
    if err := m1.Lock(context.TODO()); err != nil {
        log.Fatal(err)
    }
    fmt.Println("acquired lock for s1")

    m2Locked := make(chan struct{})
    go func() {
        defer close(m2Locked)
        // wait until s1 is locks /my-lock/
        if err := m2.Lock(context.TODO()); err != nil {
            log.Fatal(err)
        }
    }()

    if err := m1.Unlock(context.TODO()); err != nil {
        log.Fatal(err)
    }
    fmt.Println("released lock for s1")

    <-m2Locked
    fmt.Println("acquired lock for s2")

    // Output:
    // acquired lock for s1
    // released lock for s1
    // acquired lock for s2
}