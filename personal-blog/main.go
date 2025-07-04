package main

import (
	"fmt"
	"net/http"
	"os"
	"personal-blog/personal-blog/handlers"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Render(w, r, "guest", "articles", nil)
	})

	port := os.Getenv("PORT")
	fmt.Println("Server is running: http://localhost:" + port)
	http.ListenAndServe(":"+port, mux)
}
