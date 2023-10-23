package api

import (
	"net/http"
	"time"
)

func initializeServer() *http.Server {
	http.HandleFunc("/", RouteHandler)

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5,
		WriteTimeout: 5,
	}

	return server
}

func Start(flag *bool) {
	server := initializeServer()

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()

	for !*flag {
		time.Sleep(1 * time.Second)
	}

	err := server.Close()
	if err != nil {
		return
	}
}
