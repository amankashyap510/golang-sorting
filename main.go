// main.go
package main

import (
	"net/http"
	"sortmodule/handlers"
)

func main() {
	http.HandleFunc("/process-single", handlers.ProcessSingleHandler)
	http.HandleFunc("/process-concurrent", handlers.ProcessConcurrentHandler)

	port := ":8000"
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
