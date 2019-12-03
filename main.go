package main

import (
    "fmt"
	"strconv"
	"sync"
    "time"
    "github.com/go-redis/redis/v7"
)

func workerWrite(id int, wg *sync.WaitGroup, client *redis.Client) {
    defer wg.Done()
    fmt.Printf("Worker write %d starting\n", id)

	err := client.Set("key"+strconv.Itoa(id), "value"+strconv.Itoa(id), time.Minute * 3).Err()
	if err != nil {
		panic(err)
	}

    time.Sleep(time.Second)
    fmt.Printf("Worker write %d done\n", id)

}

func workerRead(id int, wg *sync.WaitGroup, client *redis.Client) {
	defer wg.Done()
	fmt.Printf("Worker read %d starting\n", id)

	val2, err := client.Get("key"+strconv.Itoa(id)).Result()
	client.Del("key"+strconv.Itoa(id))
	if err == redis.Nil {
		fmt.Println("key"+strconv.Itoa(id)+" does not exist")
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("key"+strconv.Itoa(id), val2)
	}

	time.Sleep(time.Second)
	fmt.Printf("worker read"+strconv.Itoa(id)+" done\n")

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
  	var wt sync.WaitGroup

  	for i := 1; i <=100000; i++ {
   	   wg.Add(1)
   	   go workerWrite(i, &wg, client)
  	}

  	wg.Wait()

	for i := 1; i <=100000; i++ {
		wt.Add(1)
		go workerRead(i, &wt, client)
	}

	wt.Wait()
}
