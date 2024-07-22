package proc

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var done = make(chan struct{})

func init() {
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM, syscall.SIGINT)

		for {
			v := <-signals
			switch v {
			case syscall.SIGTERM:
				stopOnSignal()
				gracefulStop(signals, syscall.SIGTERM)
			case syscall.SIGINT:
				stopOnSignal()
				gracefulStop(signals, syscall.SIGINT)
			default:
				fmt.Println("Got unregistered signal:", v)
			}
		}
	}()
}

// Done returns the channel that notifies the process quitting.
func Done() <-chan struct{} {
	return done
}
func stopOnSignal() {
	select {
	case <-done:
		// already closed
	default:
		close(done)
	}
}
