package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func main() {
	if len(os.Getenv("REDIS_HOST")) > 0 {
		ConnectRedis()
	}
	if len(os.Getenv("AMQP_HOST")) > 0 {
		ConnectRabbitMQ()
	}

	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	log.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(struct {
		DummyText    string
		RedisHost    string
		RabbitmqHost string
	}{
		DummyText:    "Hello World!",
		RedisHost:    os.Getenv("REDIS_HOST"),
		RabbitmqHost: os.Getenv("AMQP_HOST"),
	})
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

	for {
		ch, err := conn.Channel()
		if err != nil {
			panic(err)
		}

		err = ch.ExchangeDeclare(
			"test_exchange",
			"fanout",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			panic(err)
		}
		err = ch.Publish(
			"test_exchange",
			"",
			true,
			true,
			amqp.Publishing{
				Body: []byte("Hello World!"),
			},
		)
		if err != nil {
			panic(err)
		}
		log.Println("Published message")

		time.Sleep(1 * time.Second)
	}
}
