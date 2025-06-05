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
	var connBecameActive sync.WaitGroup
	connBecameActive.Add(1)

	srv := &http.Server{
		ConnState: func(c net.Conn, state http.ConnState) {
			if state == http.StateActive {
				connBecameActive.Done() // соединение активно
			}
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello!")
		}),
	}

	srv.RegisterOnShutdown(func() {
		fmt.Println("SHUTDOWN CALLBACK")
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

	<-started

	// Выполни запрос
	go func() {
		time.Sleep(300 * time.Millisecond)
		resp, err := http.Get("http://localhost:7070")
		if err == nil {
			resp.Body.Close()
		}
	}()

	// Жди, пока соединение станет активным
	connBecameActive.Wait()

	time.Sleep(300 * time.Millisecond)
	_ = srv.Shutdown(context.Background())
	wg.Wait()
	fmt.Println("Server exited cleanly.")
}
