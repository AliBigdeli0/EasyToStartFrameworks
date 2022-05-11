package pkg

import (
	"errors"
	"path"

	flags "github.com/jessevdk/go-flags"
)

var (
	err error
)

type IArgs interface {
	GetConfigPath() string
	GetBasePath() string
	checkArgumentValidation() error
}

type Args struct {
	ConfigurePath string `short:"c" long:"config" description:"configure path" required:"true"`
	BasePath      string `short:"b" long:"base" description:"base path used for reading configure files. Note: this must be absolute path" required:"true"`
}

func ArgParse(os_args []string) (IArgs, error) {
	var opts Args = Args{
		ConfigurePath: "",
		BasePath:      "",
	}
	_, err = flags.ParseArgs(&opts, os_args)
	if err != nil {
		return nil, err
	}

	err = opts.checkArgumentValidation()
	if err != nil {
		return nil, err
	}

	return &opts, nil
}

func (a *Args) checkArgumentValidation() error {

	is_exist, err := IsDirExist(a.BasePath)
	if err != nil || !is_exist {
		return errors.New("base path dose not exist")
	}

	is_exist, err = IsFileExist(path.Join(a.BasePath, a.ConfigurePath))

	if err != nil || !is_exist {
		return errors.New("invalid configuration path. The path dose not exist")
	}

	return nil
}

func (a *Args) GetConfigPath() string {
	return a.ConfigurePath
}

func (a *Args) GetBasePath() string {
	return a.BasePath
}
