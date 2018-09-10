package main

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const commandName = "producer_api"

func main() {
	godotenv.Load()
	hystrix.ConfigureCommand(
		commandName,
		hystrix.CommandConfig{
			Timeout:                500,
			MaxConcurrentRequests:  100,
			ErrorPercentThreshold:  50,
			RequestVolumeThreshold: 3,
			SleepWindow:            10000,
		},
	)

	http.HandleFunc("/", logger(handle))
	log.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	output := make(chan bool, 1)
	errors := hystrix.Go(commandName, func() error {
		err := callChargeProducerAPI()
		if err == nil {
			output <- true
		}
		return err
	}, nil)

	select {

	case out := <-output:
		log.Printf("success %v", out)
	case err := <-errors:
		log.Printf("failed %s", err)
	}
}

func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		fn(w, r)
	}
}

func callChargeProducerAPI() error {
	fmt.Printf("SERVER_ERROR %s\n", os.Getenv("SERVER_ERROR"))
	if os.Getenv("SERVER_ERROR") == "1" {
		return errors.New("503 error")
	}
	return nil
}
