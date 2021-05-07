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
	srv2 := http.Server{
		Addr: ":8081",
	}
	c := make(chan os.Signal, 1)
	shutdown := make(chan struct{}, 1)
	signal.Notify(c, os.Interrupt)
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		log.Printf("starting srv1 on %s\n", srv1.Addr)
		if err:= srv1.ListenAndServe(); !errors.Is(err, http.ErrServerClosed){
			return fmt.Errorf("srv1:%w", err)
		}
		return nil
	})

	g.Go(func() error {
		log.Printf("starting srv2 on %s\n", srv2.Addr)
		if err:= srv2.ListenAndServe(); !errors.Is(err, http.ErrServerClosed){
			return fmt.Errorf("srv2:%w", err)
		}
		return nil
	})

	g.Go(func() error {
		select {
		case sig := <-c:
			log.Printf("signal received: %s\n", sig)
			close(shutdown)
		case <-shutdown:
		}
		return nil
	})

	g.Go(func() error {
		<-shutdown
		log.Printf("closing srv1\n")
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := srv1.Shutdown(ctx); err != nil {
			srv1.Close()
			return fmt.Errorf("srv1: cloud not stop server gracefully: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		<-shutdown
		log.Printf("closing srv2\n")
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := srv2.Shutdown(ctx); err != nil {
			srv2.Close()
			return fmt.Errorf("srv2: cloud not stop server gracefully: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		_, notClosed := <- shutdown
		if notClosed{
			close(shutdown)
		}
		log.Printf("%s\n", err.Error())
		return
	}
}
