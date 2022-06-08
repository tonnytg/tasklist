package daemon

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	log.Println("Start Daemon")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	processId := os.Getpid()
	fmt.Println("processId:", processId)
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
