package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	wc, err := w.Write([]byte("Hello API\n"))

	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}

	fmt.Println(wc)
}
