package main

import (
	"api/internal/config"
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	key := make([]byte, 64)
	fmt.Println(key)
}

func main() {
	config.Load()
	r := router.Generate()
	fmt.Printf("listening on port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
