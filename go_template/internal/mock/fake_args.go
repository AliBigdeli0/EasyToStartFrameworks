package mock

import (
	"fmt"
	"os"
)

func ValidConfigureArgs() []string {
	wd, _ := os.Getwd()
	wd_string := "--config=config.yml"
	bd_string := fmt.Sprintf("--base=%s/../../bin/", wd)

	return []string{
		wd_string,
		bd_string,
	}
}
