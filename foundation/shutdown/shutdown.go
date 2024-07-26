package foundation

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// InitGracefulShutdownHandling generates the channel that listens to SIGTERM and a waitgroup to use for finishing work when shutting down
func InitGracefulShutdownHandling() (gracefulShutdown chan os.Signal, waitGroup *sync.WaitGroup) {

	// define channel used to trigger graceful shutdown
	gracefulShutdown = make(chan os.Signal)

	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)

	waitGroup = &sync.WaitGroup{}

	return gracefulShutdown, waitGroup
}

// HandleGracefulShutdown waits for SIGTERM to unblock gracefulShutdown and waits for the waitgroup to await pending work
func HandleGracefulShutdown(gracefulShutdown chan os.Signal, waitGroup *sync.WaitGroup, functionsOnShutdown ...func()) {

	signalReceived := <-gracefulShutdown
	log.Info().
		Msgf("Received signal %v. Waiting for running tasks to finish...", signalReceived)

	// execute any passed function
	for _, f := range functionsOnShutdown {
		f()
	}

	waitGroup.Wait()

	log.Info().Msg("Shutting down...")
}

// InitCancellationContext adds cancelation to a context and on sigterm triggers the cancel function
func InitCancellationContext(ctx context.Context) context.Context {

	ctx, cancel := context.WithCancel(context.Background())

	// define channel used to trigger cancellation
	cancelChannel := make(chan os.Signal)

	signal.Notify(cancelChannel, syscall.SIGTERM, syscall.SIGINT)

	go func(cancelChannel chan os.Signal, cancel context.CancelFunc) {
		<-cancelChannel
		cancel()
	}(cancelChannel, cancel)

	return ctx
}
