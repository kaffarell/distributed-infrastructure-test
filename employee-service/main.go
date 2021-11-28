package main
import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
	"encoding/json"
)

type Employee struct {
	Name string `json:"name"`
	Address string `json:"address"`
}

func main() {
	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	json, err := json.Marshal(Employee{Name: "Gordon", Address: "San Jose"})
	if err != nil {
		fmt.Println("Error converting object to json: ")
		fmt.Println(err)
	}
	fmt.Println(string(json))

	err = client.Set(ctx, "add123", json, 0).Err()
	if err != nil {
		fmt.Println("Error setting key: ")
		fmt.Println(err)
	}

	val, err := client.Get(ctx, "add123").Result()
	if err != nil {
		fmt.Println("Error getting key: ")
		fmt.Println(err)
	}
	fmt.Println("Value got from redis:")
	fmt.Println(val)
}