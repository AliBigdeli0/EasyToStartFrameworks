package worker

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/avi_project/internal/pkg"
)

type SignalWorker struct {
	context           context.Context
	channel           chan Job
	terminate_channel chan os.Signal
	cancelFun         context.CancelFunc
}

func NewSignalWorker(ctx context.Context, cancelFunc context.CancelFunc) Worker {
	return &SignalWorker{
		context:           ctx,
		channel:           make(chan Job),
		terminate_channel: make(chan os.Signal),
		cancelFun:         cancelFunc,
	}
}

func (s *SignalWorker) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	go s.ChannelActions()

	signal.Notify(s.terminate_channel, syscall.SIGINT, syscall.SIGTERM)

	<-s.terminate_channel

	s.cancelFun()
	pkg.LogInst().Info(fmt.Sprintf("DANGER: going to stop '%s'", s.GetName()))
}

func (s *SignalWorker) Close() {

}

func (s *SignalWorker) ChannelActions() {
	for {
		select {
		case job := <-s.channel:
			if job.Id == KILL_ALL {
				s.Close()
			}
		case <-s.context.Done():
			s.Close()
		}
	}
}

func (s *SignalWorker) GetName() string {
	return "signal_worker"
}

func (s *SignalWorker) GetInChannel() chan Job {
	return s.channel
}
