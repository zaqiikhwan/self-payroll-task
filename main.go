package main

import (
	"self-payrol/config"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Infof(".env is not loaded properly")
	}
	log.Infof("read .env from file")

	// fmt.Println(os.Getenv("DATABASE_URL"))

	config := config.NewConfig()
	// fmt.Println(config)
	server := InitServer(config)
	// server.Run()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Run()
	}()

	wg.Wait()
}
