package bruteforce

import (
    "sync"
    "github.com/panjf2000/ants/v2"
)

type WorkerPool struct {
    pool    *ants.Pool
    wg      sync.WaitGroup
}

func NewWorkerPool(size int) (*WorkerPool, error) {
    pool, err := ants.NewPool(size)
    if err != nil {
        return nil, err
    }
    return &WorkerPool{pool: pool}, nil
}

func (wp *WorkerPool) Submit(task func()) error {
    wp.wg.Add(1)
    return wp.pool.Submit(func() {
        defer wp.wg.Done()
        task()
    })
}

func (wp *WorkerPool) Wait() {
    wp.wg.Wait()
}

func (wp *WorkerPool) Release() {
    wp.pool.Release()
}
