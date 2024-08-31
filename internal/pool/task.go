package pool

import (
    "errors"
    "log"
    "time"

    "github.com/go-redis/redis/v8"
    "context"
)

type Task struct {
    ID      string
    Payload string
    Retries int
}

func (t *Task) Process(redisAddr string) error {
    log.Printf("Task ID: %s ishlamoqda...", t.ID)

    rdb := redis.NewClient(&redis.Options{
        Addr: redisAddr,
    })
    ctx := context.Background()

    if err := someProcessing(t.Payload); err != nil {
        log.Printf("Task ID: %s xato: %v", t.ID, err)
        rdb.Set(ctx, "task:"+t.ID+":status", "failed", 0)
        return err
    }

    rdb.Set(ctx, "task:"+t.ID+":status", "completed", 0)
    return nil
}

func someProcessing(payload string) error {
    if payload == "fail" {
        return errors.New("vazifa bajarilmadi")
    }
    time.Sleep(2 * time.Second)
    return nil
}
