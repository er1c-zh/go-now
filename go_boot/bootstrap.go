package go_boot

import (
	"github.com/er1c-zh/go-now/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type GracefulExitHandler func()

var (
	exitHandlers []GracefulExitHandler

	activeExitChanOnce sync.Once
	activeExitChan     chan struct{}

	systemSignalChan chan os.Signal
	exitSignalSet    map[syscall.Signal]struct{}
)

func init() {
	activeExitChan = make(chan struct{})
	systemSignalChan = make(chan os.Signal)

	exitSignalSet = map[syscall.Signal]struct{}{
		syscall.SIGHUP:  {},
		syscall.SIGINT:  {},
		syscall.SIGQUIT: {},
		syscall.SIGTERM: {},
		syscall.SIGKILL: {},
	}
}

/*
注册退出事件的函数
*/
func RegisterExitHandlers(h GracefulExitHandler) {
	exitHandlers = append(exitHandlers, h)
}

/*
主动退出
*/
func Exit() {
	activeExitChanOnce.Do(func() {
		activeExitChan <- struct{}{}
	})
}

/*
阻塞到退出
依次调用注册的退出响应函数
有超时时间，传0为无超时时间
*/
func WaitExit(expireTime time.Duration) {
	signal.Notify(systemSignalChan, syscall.SIGTERM, syscall.SIGINT)

	toExit := false
	for !toExit {
		select {
		case <-activeExitChan:
			toExit = true
		case sig := <-systemSignalChan:
			afterConvert, ok := sig.(syscall.Signal)
			if !ok {
				continue
			}
			if _, ok := exitSignalSet[afterConvert]; !ok {
				continue
			}
			toExit = true
		}

	}

	if expireTime == 0 {
		for _, f := range exitHandlers {
			f()
		}
		log.Info("Exit gracefully!")
		return
	}

	t := time.NewTimer(expireTime)
	finishExitChan := make(chan struct{})

	go func() {
		for _, f := range exitHandlers {
			f()
		}
		finishExitChan <- struct{}{}
	}()

	select {
	case <-finishExitChan:
		log.Info("Exit gracefully!")
	case <-t.C:
		log.Warn("Exit handler process timeout.")
	}
}

