package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := r.Body.Close(); err != nil {
				fmt.Println("failed to handle '/' request:", err)
			}
		}()

		b, err := os.ReadFile("./static/index.html")
		if err != nil {
			w.Write([]byte("<html><header><title>chirpr</title></header><body>Failed loading index.html</body></html>"))
			return
		}

		_, err = w.Write(b)
		fmt.Println(err)
	})

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("failed to server:", err)
	}
}
