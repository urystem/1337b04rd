package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Создание клиента Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // адрес Redis
		Password: "",               // пароль, если есть
		DB:       0,                // номер БД (0 по умолчанию)
	})

	// Пример: Установить значение с TTL
	err := rdb.Set(ctx, "mykey", "hello redis", 10*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// Пример: Получить значение
	val, err := rdb.Get(ctx, "mykey").Result()
	if err == redis.Nil {
		fmt.Println("Ключ не найден")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("mykey:", val)
	}

	// Пример: Проверка TTL
	ttl, err := rdb.TTL(ctx, "mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("TTL:", ttl)
}
