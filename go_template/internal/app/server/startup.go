package server

import (
	"os"

	"github.com/avi_project/internal/pkg"
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

}

func Run() {
	init_items()
	defer close()
	pkg.LogInst().Info("server started...")

	pkg.LogInst().Info("server is going to shutdown")
}
