package test

import (
	"strings"
	"testing"

	mock "github.com/avi_project/internal/mock"
	"github.com/avi_project/internal/pkg"
)

func TestConfigPath(t *testing.T) {

	valid_args := mock.ValidConfigureArgs()

	args, _ := pkg.ArgParse(valid_args)

	path := pkg.GetConfigPath(&args)
	if path == "" {
		t.Fail()
	}

	if !strings.HasSuffix(path, args.GetConfigPath()) {
		t.Fail()
	}
}
