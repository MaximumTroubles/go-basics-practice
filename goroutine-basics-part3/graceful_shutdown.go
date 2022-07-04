package main

import (
	"os"
	"os/signal"
)

func main() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// run main process/jobs

	<-signalCh
	// stop logic
}

// it siplifised code whats going on inside Notify

func Notify(signalCh) {
	//..
	select {
	case signalCh <- sig:
	default:
	}
	//..
}
