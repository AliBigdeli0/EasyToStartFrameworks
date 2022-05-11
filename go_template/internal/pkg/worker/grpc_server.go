package worker

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/avi_project/internal/pkg"
	"google.golang.org/grpc"
)

var (
	grpc_server *grpc.Server
)

type GrpcServerWorker struct {
	ListenPort int
	context    context.Context
	channel    chan Job
}

func NewGrpcServer(ctx context.Context, config pkg.IConfig) Worker {
	return &GrpcServerWorker{
		ListenPort: 81,
		context:    ctx,
		channel:    make(chan Job),
	}
}

func (s *GrpcServerWorker) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	go s.ChannelActions()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.ListenPort))
	if err != nil {
		pkg.LogInst().Error(err.Error())
	}

	grpc_server = grpc.NewServer()

	pkg.LogInst().Info(fmt.Sprintf("going to start grpc server at port: %d", s.ListenPort))

	if err = grpc_server.Serve(lis); err != nil {
		pkg.LogInst().Error(err.Error())
	} else {
		pkg.LogInst().Info("grpc server started...")
	}

	pkg.LogInst().Info(fmt.Sprintf("DANGER: going to stop grpc server %s", s.GetName()))
}

func (s *GrpcServerWorker) Close() {
	grpc_server.Stop()
}

func (w *GrpcServerWorker) GetInChannel() chan Job {
	return w.channel
}

func (w *GrpcServerWorker) GetName() string {
	return "grpc_worker"
}

func (w *GrpcServerWorker) ChannelActions() {
	for {
		select {
		case job := <-w.channel:
			if job.Id == KILL_ALL {
				w.Close()
			}
		case <-w.context.Done():
			w.Close()
		}

	}
}
