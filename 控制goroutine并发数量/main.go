/*
通过sync.WaitGroup + channel控制数量
参考k8s.io/client-go/util/workqueue/parallelizer.go
*/

package main

import (
	"context"
	"sync"
)

type DoWorkPieceFunc func(piece int)

// Parallelize is a very simple framework that allows for parallelizing
// N independent pieces of work.
//
// Deprecated: Use ParallelizeUntil instead.
func Parallelize(workers, pieces int, doWorkPiece DoWorkPieceFunc) {
	ParallelizeUntil(nil, workers, pieces, doWorkPiece)
}

// ParallelizeUntil is a framework that allows for parallelizing N
// independent pieces of work until done or the context is canceled.
func ParallelizeUntil(ctx context.Context, workers, pieces int, doWorkPiece DoWorkPieceFunc) {
	var stop <-chan struct{}
	if ctx != nil {
		stop = ctx.Done()
	}

	toProcess := make(chan int, pieces)
	for i := 0; i < pieces; i++ {
		toProcess <- i
	}
    //不影响channel的读
	close(toProcess)

	if pieces < workers {
		workers = pieces
	}

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer utilruntime.HandleCrash()
			defer wg.Done()
            //多个workers并发从toProcess中读
			for piece := range toProcess {
				select {
				case <-stop:
					return
				default:
					doWorkPiece(piece)
				}
			}
		}()
	}
	wg.Wait()
}
