package main

import (
	"encoding/json"
	"fmt"
	//Redis "github.com/darthchudi/gbedu-go/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func sayHi(w http.ResponseWriter, r *http.Request) {
	payload := map[string]interface{}{
		"name": "Fuck!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		fmt.Println("An error occured while encoding response")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", sayHi)
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3007",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//redis := Redis.CreateRedisInstance()
	log.Fatal(server.ListenAndServe())

}
