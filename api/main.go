package main

import (
	router "api/src/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("init")

	r := router.Generate()
	fmt.Println(r)

	log.Fatal(http.ListenAndServe(":5000", r))
}
