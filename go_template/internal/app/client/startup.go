package client

import (
	"context"
	"os"

	pkg "github.com/avi_project/internal/pkg"
	"github.com/avi_project/internal/pkg/worker"
)

var (
	err error
)

func init_items() {
	var arguments pkg.IArgs
	sys_args := os.Args

	arguments, err = pkg.ArgParse(sys_args)

	if err != nil {
		panic(err)
	}

	err = pkg.ConfigInit(&arguments)
	if err != nil {
		panic(err)
	}

	pkg.LogInit(pkg.ConfigInst(), &arguments)
}

func close() {
	pkg.LogInst().Info("DANGER: going to stop hole program")
}

func Run() {
	init_items()
	defer close()
	pkg.LogInst().Info("client started...")

	ctx, cancelFunc := context.WithCancel(context.Background())
	worker.InitServerWorkers(ctx, cancelFunc)
	worker.RunAllWorkers()
}
