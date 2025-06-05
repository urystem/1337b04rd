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
		time.Sleep(10 * time.Second)
	})

	var wg sync.WaitGroup
	wg.Add(1)

	started := make(chan struct{})

	go func() {
		defer wg.Done()
		ln, err := net.Listen("tcp", ":7070")
		if err != nil {
			fmt.Println("LISTEN ERROR:", err)
			return
		}
		fmt.Println("Server started on :7070")
		close(started)
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			fmt.Println("SERVER ERROR:", err)
		}
	}()

	<-started // жди запуска сервера

	// Явно выполни запрос, чтобы точно активировать соединение
	time.Sleep(300 * time.Millisecond) // небольшая задержка
	resp, err := http.Get("http://localhost:7070")
	if err != nil {
		fmt.Println("REQUEST ERROR:", err)
	} else {
		resp.Body.Close()
	}

	// Подожди чуть-чуть, чтобы сервер успел обработать
	time.Sleep(300 * time.Millisecond)

	// Заверши сервер
	err = srv.Shutdown(context.Background())
	if err != nil {
		fmt.Println("SHUTDOWN ERROR:", err)
	}

	wg.Wait()
	fmt.Println("Server exited cleanly.")
}
