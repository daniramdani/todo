package main

import (
	"log"
	"os"
	"sync"
	"github.com/daniramdani/todo/router"
	config "github.com/joho/godotenv"
)

func main() {
	// read env 
	err := config.Load(".env")
	if err != nil {
		log.Printf("can't load .env file")
		os.Exit(2)
	}
	cfgenv := os.Getenv("ENV")
	log.Printf("environment ENV=%s", cfgenv)

	//oConfig.InitPanicPublisher()
	service := router.MakeHandler()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		service.HTTPServerMain()
	}()

	// Wait All services to end
	wg.Wait()
}
