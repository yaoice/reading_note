/*
定义一个包含Error的struct
再定义一个无缓冲的channel实现同步传递
 */
package main

type Result struct {
    Error error
}

func updateNode(node string) error {
    return nil
}

func main() {
    nodeList := make([]string, 10)
    checkStatus := func(done <-chan interface{}, nodeList ...string) <-chan Result {
        results := make(chan Result)
        go func() {
            defer close(results)
            for _, node := range nodeList {
                var result Result
                // update node
                err := updateNode(node)
                result = Result{Error: err}
                select {
                case <- done:
                    return
                case results <- result:
                }
            }
        }()
        return results
    }

    done := make(chan interface{})
    defer close(done)

    for result := range checkStatus(done, nodeList...) {
        if result.Error != nil {
            panic(result.Error)
        }
    }
}
