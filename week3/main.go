package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	srv1 := http.Server{
		Addr: ":8080",
	}
	c := make(chan os.Signal, 1)
	shutdown := make(chan struct{},1)
	signal.Notify(c, os.Interrupt)
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		log.Printf("starting srv1 on %s\n", srv1.Addr)
		return fmt.Errorf("srv1:%w", srv1.ListenAndServe())
	})

	g.Go(func() error {
		select {
			case sig:=<-c:
				log.Printf("closing srv1: %s\n", sig)
			case <-shutdown:
		}
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err:= srv1.Shutdown(ctx); err!=nil{
			srv1.Close()
			return fmt.Errorf("srv1: cloud not stop server gracefully: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		close(shutdown)
		log.Printf("%s\n", err.Error())
		return
	}
}
