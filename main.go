package main

import (
	"context"
	"flag"
	"log"
	"github.com/crossle/hacker-news-mixin-bot/durable"
	"github.com/crossle/hacker-news-mixin-bot/services"
)

func main() {
	service := flag.String("service", "blaze", "run a service")
	flag.Parse()
	// 解析 CMD 命令
	db := durable.OpenDatabaseClient(context.Background())
	defer db.Close()

	switch *service {
	case "blaze":
		err := StartBlaze(db) // start up new ws server conn
		if err != nil {
			log.Println(err)
		}
	default:
		hub := services.NewHub(db) // new main service obj
		err := hub.StartService(*service)
		if err != nil {
			log.Println(err)
		}
	}
}
