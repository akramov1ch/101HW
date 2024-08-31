package main

import (
	"101HW/config"
	"101HW/internal/http"
	"101HW/internal/pool"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Konfiguratsiyani yuklashda xatolik: %v", err)
	}

	taskQueue := make(chan pool.Task, config.QueueSize)
	wp := pool.NewWorkerPool(config.WorkerCount, taskQueue, config.RedisAddr)
	wp.Start()

	server := http.NewServer(config.ServerAddr, wp)
	log.Printf("Server %s portida ishga tushdi", config.ServerAddr)
	log.Fatal(server.ListenAndServe())
}
