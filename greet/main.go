package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	response := fmt.Sprintf("Hello %s", name)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	port := 9090

	http.HandleFunc("/greet", handleGreet)

	fmt.Printf("Server started on port %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
