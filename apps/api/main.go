package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHello)
	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	wc, err := w.Write([]byte("Hello API\n"))

	if err != nil {
		slog.Error("error writing response:", err)
		return
	}

	fmt.Println(wc)
}
