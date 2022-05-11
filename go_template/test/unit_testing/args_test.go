package test

import (
	"testing"

	mock "github.com/avi_project/internal/mock"
	"github.com/avi_project/internal/pkg"
	"github.com/stretchr/testify/assert"
)

func TestArgsParser(t *testing.T) {
	valid_args := mock.ValidConfigureArgs()
	println(valid_args[0])
	args, err := pkg.ArgParse(valid_args)
	if err != nil {
		assert.Fail(t, "args has errors: ", err)
	}

	_, err = pkg.ArgParse([]string{})
	if err == nil {
		assert.Fail(t, "arg parse can not be empty")
	}

	println(args.GetBasePath())

	println(args.GetConfigPath())
}
