package server

import (
	"fmt"
	"lec-02/handlers"
	"net/http"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8090")

	defer func() {
		fmt.Println("Server is stopped")
	}()

	http.HandleFunc("/sum", handlers.CalSum)
	http.HandleFunc("/sub", handlers.CalSub)
	http.HandleFunc("/mul", handlers.CalMul)
	http.HandleFunc("/div", handlers.CalDiv)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic("Error when running server")
	}
}
