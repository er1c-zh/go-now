package go_boot

import (
	"github.com/er1c-zh/go-now/log"
	"testing"
)

func TestDemo(t *testing.T) {
	defer log.Flush()
	RegisterExitHandlers(func() {
		log.Info("handlers 1")
	})
	RegisterExitHandlers(func() {
		log.Info("handlers 2")
	})
	RegisterExitHandlers(func() {
		log.Info("handlers 3")
	})

	go func() {
		Exit()
	}()

	WaitExit(0)
}
