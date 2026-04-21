package core_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envconfig:"ADDR" reqired:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" reqired:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("HTTP", &config); err != nil {
		return Config{}, fmt.Errorf("proccess envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	conf, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get Config: %w", err)
		panic(err)
	}
	return conf
}
