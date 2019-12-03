/*
定义一个error类型的Slice, 通过加读写锁对它进行操作

选自：kubernetes/pkg/scheduler/core/generic_scheduler.go的代码片段
 */
package main

func main() {
    /*
    var (
        mu   = sync.Mutex{}
        wg   = sync.WaitGroup{}
        errs []error
    )
    appendError := func(err error) {
        mu.Lock()
        defer mu.Unlock()
        errs = append(errs, err)
    }

    results := make([]schedulerapi.HostPriorityList, len(priorityConfigs), len(priorityConfigs))

    // DEPRECATED: we can remove this when all priorityConfigs implement the
    // Map-Reduce pattern.
    for i := range priorityConfigs {
        if priorityConfigs[i].Function != nil {
            wg.Add(1)
            go func(index int) {
                defer wg.Done()
                var err error
                results[index], err = priorityConfigs[index].Function(pod, nodeNameToInfo, nodes)
                if err != nil {
                    appendError(err)
                }
            }(i)
        } else {
            results[i] = make(schedulerapi.HostPriorityList, len(nodes))
        }
    }

    // Wait for all computations to be finished.
    wg.Wait()
    if len(errs) != 0 {
        return schedulerapi.HostPriorityList{}, errors.NewAggregate(errs)
    }
    */
}
