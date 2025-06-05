package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	srv := &http.Server{Addr: ":7070"}

	srv.RegisterOnShutdown(func() {
		fmt.Println("SHUTDOWN CALLBACK")
	})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println("SERVER ERROR:", err)
		}
	}()
	time.Sleep(4 * time.Second)
	srv.Shutdown(context.Background())
	wg.Wait()
	fmt.Println("dd")
}

/*
student@ALEM-F4-R05-06:~/my/1337b04rd-my$ go run server/main.go
SERVER ERROR: http: Server closed
dd
*/

//регистр кейде шақырылмайды
//оған себеп