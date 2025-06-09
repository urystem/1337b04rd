package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	srv := &http.Server{Addr: ":7070"}

	wg.Add(2) // например, у нас 2 задачи на завершение

	srv.RegisterOnShutdown(func() {
		fmt.Println("Shutdown task 1 started")
		time.Sleep(2 * time.Second) // имитация работы
		fmt.Println("Shutdown task 1 done")
		wg.Done()
	})

	srv.RegisterOnShutdown(func() {
		fmt.Println("Shutdown task 2 started")
		time.Sleep(1 * time.Second)
		fmt.Println("Shutdown task 2 done")
		wg.Done()
	})

	go func() {
		// if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	time.Sleep(3 * time.Second) // ждем, чтобы сервер "поработал"

	// fmt.Println("Calling shutdown...")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// cancel()

	if err := srv.Close(); err != nil {
		fmt.Println("Shutdown error:", err)
	}

	wg.Wait() // ждём, пока все shutdown-функции завершатся

	fmt.Println("All shutdown tasks done, exiting")
}
