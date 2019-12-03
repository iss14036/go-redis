package main

import (
    "fmt"
	"strconv"
	"sync"
    "time"
    "github.com/go-redis/redis/v7"
)

func worker(id int, wg *sync.WaitGroup, client *redis.Client) {
    fmt.Printf("Worker %d starting\n", id)

	err := client.Set("key"+strconv.Itoa(id), "value"+strconv.Itoa(id), time.Minute).Err()
	if err != nil {
		panic(err)
	}

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)

    wg.Done()
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 1000,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

  	var wg sync.WaitGroup

  	for i := 1; i <=100000; i++ {
   	   wg.Add(1)
   	   worker(i, &wg, client)
  	}

  	wg.Wait()
}
