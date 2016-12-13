package main

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type (
	config struct {
		Token string
		Shards int
		Owner string
		Prefix string
	}
)

func loadConfig(file string) (*config, error) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var conf config
	_, err = toml.Decode(string(body), &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
