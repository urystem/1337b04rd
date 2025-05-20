package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background() // ok, корневой контекст

	// пока не уверен, что за контекст нужен здесь — поставим TODO:
	result := fetchData(context.TODO()) // вернёмся позже и решим
	fmt.Println(result, ctx)
}

func fetchData(context.Context) error
