package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func unmarshallPrototype(configFilePath string, prototype interface{}) (output interface{}, err error) {
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading the config file %v", err)
	}

	err = yaml.Unmarshal(bytes, prototype)
	if err != nil {
		return nil, fmt.Errorf("error reading the config file %v", err)
	}

	return prototype, nil
}
