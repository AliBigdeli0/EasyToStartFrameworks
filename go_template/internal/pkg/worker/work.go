package worker

import (
	"context"
	"sync"

	"github.com/avi_project/internal/pkg"
)

var (
	_wg             *sync.WaitGroup
	wks             map[string]Worker
	workers_channel chan Job
)

type Worker interface {
	Run(wg *sync.WaitGroup)
	ChannelActions()

	//properties
	GetName() string
	GetInChannel() chan Job
}

func InitServerWorkers(ctx context.Context, cf context.CancelFunc) {

	workers_channel = make(chan Job)

	wks = make(map[string]Worker)

	//init grpc
	grpc_worker := NewGrpcServer(ctx, pkg.ConfigInst())
	wks[grpc_worker.GetName()] = grpc_worker

	//init signal
	signal_worker := NewSignalWorker(ctx, cf)
	wks[signal_worker.GetName()] = signal_worker

	go handleChannels()
}

func RunAllWorkers() {
	_wg = &sync.WaitGroup{}
	_wg.Add(len(wks))
	for k := range wks {
		go wks[k].Run(_wg)
	}
	_wg.Wait()
}

func GetWorkersChannel() chan Job {
	return workers_channel
}

func handleChannels() {
	for job := range workers_channel {
		if job.Id == KILL_ALL {
			for k := range wks {
				wks[k].GetInChannel() <- job
			}
		}
	}
}
