package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

func main() {
	srv := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello!")
		}),
	}

	srv.RegisterOnShutdown(func() {
		fmt.Println("SHUTDOWN CALLBACK")
	})

	var wg sync.WaitGroup
	wg.Add(1)

	// Канал для синхронизации старта
	started := make(chan struct{})

	go func() {
		defer wg.Done()
		ln, err := net.Listen("tcp", ":7070")
		if err != nil {
			fmt.Println("LISTEN ERROR:", err)
			return
		}
		fmt.Println("Server started on :7070")
		close(started) // сервер успешно запущен

		err = srv.Serve(ln)
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("SERVER ERROR:", err)
		}
	}()

	// Дождись запуска сервера
	<-started

	// Сделай тестовый запрос, чтобы точно было соединение
	go func() {
		time.Sleep(500 * time.Millisecond)
		http.Get("http://localhost:7070")
	}()

	// Подожди немного перед завершением
	time.Sleep(2 * time.Second)

	// Заверши сервер
	err := srv.Shutdown(context.Background())
	if err != nil {
		fmt.Println("SHUTDOWN ERROR:", err)
	}

	wg.Wait()
	fmt.Println("Server exited cleanly.")
}
