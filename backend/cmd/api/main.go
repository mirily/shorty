package main

import (
	"fmt"
	"net/http"

	apphttp "github.com/mirily/shorty/internal/http"
)

func main() {
	router := apphttp.NewRouter()

	fmt.Println("Shorty API started on :8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}
