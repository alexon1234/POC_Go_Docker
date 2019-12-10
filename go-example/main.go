package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Need to set up enviroment variables")
	}

	ConnectRedis()
	ConnectRabbitMQ()

	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	log.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func ConnectRedis() {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Println("Redis connected")
}

func ConnectRabbitMQ() {
	conn, err := amqp.Dial(os.Getenv("AMQP_HOST"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("RabbitMQ connected")
}
