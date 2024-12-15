package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("api on\n")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
