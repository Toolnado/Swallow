package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Toolnado/SwalloW/config"
	"github.com/Toolnado/SwalloW/internal/handler"
	"github.com/Toolnado/SwalloW/internal/repository"
	"github.com/Toolnado/SwalloW/internal/service"
	"github.com/Toolnado/SwalloW/server"
)

func main() {
	config := config.NewConfig()
	db, err := repository.NewPostgresDB(config)

	if err != nil {
		log.Printf("[error conection to database: %s]\n", err)
	}

	repo := repository.NewRepository(db)

	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	server := server.NewServer(config.ServerPort, handler.InitRouters())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go server.Run()

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	server.Stop(ctx)
	cancel()
}
