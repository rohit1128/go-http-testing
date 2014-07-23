package main

import (
	"fmt"
	"net/http"
	"api"
)
func main() {
	fmt.Println("Server starting")
	http.ListenAndServe(":3000", api.Handlers())
}

