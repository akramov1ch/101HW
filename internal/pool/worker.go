package pool

import (
    "log"
    "sync"
)

type WorkerPool struct {
    Workers   int
    TaskQueue chan Task
    RedisAddr string
    wg        sync.WaitGroup
}

func NewWorkerPool(workers int, taskQueue chan Task, redisAddr string) *WorkerPool {
    return &WorkerPool{
        Workers:   workers,
        TaskQueue: taskQueue,
        RedisAddr: redisAddr,
    }
}

func (wp *WorkerPool) Start() {
    for i := 0; i < wp.Workers; i++ {
        wp.wg.Add(1)
        go wp.worker(i)
    }
}

func (wp *WorkerPool) worker(id int) {
    defer wp.wg.Done()
    log.Printf("Worker %d boshladi", id)
    for task := range wp.TaskQueue {
        if err := task.Process(wp.RedisAddr); err != nil {
            log.Printf("Worker %d vazifani bajara olmadi: %v", id, err)
        } else {
            log.Printf("Worker %d vazifani bajardi", id)
        }
    }
}

func (wp *WorkerPool) Stop() {
    close(wp.TaskQueue)
    wp.wg.Wait()
}
