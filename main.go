package main

import (
	"fmt"
	"net/http"
	"os"

	"ascii-art-export/handlers"
)

var Port = "8050"

func main() {
	if len(os.Args) != 1 {
		return
	}
	fmt.Println("http://localhost:" + Port)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.OutputHandler)
	http.HandleFunc("/Download", handlers.Download_Handler)

	err := http.ListenAndServe(":"+Port, http.DefaultServeMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
