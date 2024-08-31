package http

import (
    "encoding/json"
    "net/http"
    "101HW/internal/pool"

    "github.com/go-redis/redis/v8"
    "context"
)

type Server struct {
    addr      string
    wp        *pool.WorkerPool
}

func NewServer(addr string, wp *pool.WorkerPool) *Server {
    return &Server{
        addr: addr,
        wp:   wp,
    }
}

func (s *Server) ListenAndServe() error {
    http.HandleFunc("/submit-task", s.SubmitTaskHandler)
    http.HandleFunc("/task-status/", s.TaskStatusHandler)
    return http.ListenAndServe(s.addr, nil)
}

func (s *Server) SubmitTaskHandler(w http.ResponseWriter, r *http.Request) {
    var task pool.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, "Yaroqsiz ma'lumot", http.StatusBadRequest)
        return
    }

    s.wp.TaskQueue <- task  // To'g'ri: TaskQueue dan foydalaning
    w.WriteHeader(http.StatusAccepted)
    json.NewEncoder(w).Encode(map[string]string{"status": "accepted", "task_id": task.ID})
}

func (s *Server) TaskStatusHandler(w http.ResponseWriter, r *http.Request) {
    taskID := r.URL.Path[len("/task-status/"):]
    rdb := redis.NewClient(&redis.Options{
        Addr: s.wp.RedisAddr,  // To'g'ri: RedisAddr dan foydalaning
    })
    ctx := context.Background()

    status, err := rdb.Get(ctx, "task:"+taskID+":status").Result()
    if err != nil {
        http.Error(w, "Vazifa topilmadi", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"task_id": taskID, "status": status})
}
