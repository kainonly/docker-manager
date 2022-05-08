package common

import (
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/weplanx/schedule/utiliy"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func SetValues() (values *Values, err error) {
	if _, err = os.Stat("./config/config.yml"); os.IsNotExist(err) {
		err = errors.New("静态配置不存在，请检查路径 [./config/config.yml]")
		return
	}
	var b []byte
	b, err = ioutil.ReadFile("./config/config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(b, &values)
	if err != nil {
		return
	}
	return
}

type Inject struct {
	Values   *Values
	Log      *zap.Logger
	Schedule *utiliy.Schedule
	Js       nats.JetStreamContext
	Store    nats.ObjectStore
}

type Values struct {
	Namespace string `yaml:"namespace"`
	Debug     bool   `yaml:"debug"`
	Nats      Nats   `yaml:"nats"`
}

type Nats struct {
	Hosts []string `yaml:"hosts"`
	Nkey  string   `yaml:"nkey"`
}
