package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type Config struct {
	Receivers map[string]ReceiverConfig `yaml:"receivers"`
}

type SafeConfig struct {
	sync.RWMutex
	C *Config
}

type ReceiverConfig struct {
	AccessToken string `yaml:"access_token"`
	Fsurl       string `yaml:"fsurl"`
}

func (sc *SafeConfig) ReloadConfig(configFile string) error {
	var c = &Config{}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		return err
	}

	sc.Lock()
	sc.C = c
	sc.Unlock()

	return nil
}

func (sc *SafeConfig) GetConfigByName(name string) (*ReceiverConfig, error) {
	sc.Lock()
	defer sc.Unlock()
	if receiver, ok := sc.C.Receivers[name]; ok {
		return &receiver, nil
	}
	return &ReceiverConfig{"", ""}, fmt.Errorf("no credentials found for receiver %s", name)
}
