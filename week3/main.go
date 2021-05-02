package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	srv1 := http.Server{
		Addr: ":8080",
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		log.Println("starting srv1 ")
		return srv1.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			log.Println("closing srv1 ")
			return srv1.Shutdown(ctx)
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("%s\n", err.Error())
		return
	}
}
