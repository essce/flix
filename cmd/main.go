package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/essce/flix/pkg/handler"
	"github.com/essce/flix/pkg/redis"
	"github.com/essce/flix/pkg/tvmaze"
)

func main() {
	fmt.Println("hello")

	r, err := redis.New("")
	if err != nil {
		panic(fmt.Sprintf("unable to start redis: %s", err.Error()))
	}
	defer r.Close()

	a := tvmaze.New()

	h := handler.Handler{
		API:   a,
		Cache: r,
	}

	s := http.Server{
		Addr:         ":8000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      h.HTTPHandler(),
	}

	s.ListenAndServe()
	fmt.Println("flix listening on port 8000...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-c
		fmt.Println("goodbye-1")

		return
	}()

	fmt.Println("goodbye-2")
}
