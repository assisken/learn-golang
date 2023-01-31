package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func someLock() int {
	fmt.Println("working")
	time.Sleep(10 * time.Second)
	return 1
}

func main() {
	log := log.Default()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		s := <-done
		log.Println("received signal: ", s.String())
		cancel()
	}()

	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			someLock()
		case <-ctx.Done():
			return
		}
	}
}
