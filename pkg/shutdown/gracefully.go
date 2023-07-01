package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

func Gracefully() {
	quit := make(chan os.Signal, 1)
	defer close(quit)

	// listen for interrupt signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
