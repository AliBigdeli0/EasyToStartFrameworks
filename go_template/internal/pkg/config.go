package pkg

import (
	"fmt"
	"io/ioutil"
	"path"

	yaml "gopkg.in/yaml.v3"
)

type IConfig interface {
	ConfigurePath() string
}

type Config struct {
	Common struct {
		Log struct {
			ConfigFileName string `yaml:"configFileName"`
		}
	}
	IConfig
}

var confInst *Config

func ConfigInit(args *IArgs) error {
	confInst = &Config{}
	c, e := readConfig(path.Join((*args).GetBasePath(), (*args).GetConfigPath()))
	if e != nil {
		return e
	}
	confInst = c
	return nil
}

func readConfig(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)

	if c.Common.Log.ConfigFileName == "" {
		panic("config is empty")
	}

	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func ConfigInst() IConfig {
	return confInst
}

func (c *Config) ConfigurePath() string {
	return c.Common.Log.ConfigFileName
}
