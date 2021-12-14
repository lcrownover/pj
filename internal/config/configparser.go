package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v3"
)

type ProjectConfig struct {
	Name         string   `yaml:"name"`
	Directory    string   `yaml:"directory"`
	Command      string   `yaml:"command"`
	Precommands  []string `yaml:"pre_commands"`
	Postcommands []string `yaml:"post_commands"`
	Navigate     bool     `yaml:"navigate"`
}

type ProjectList struct {
	Projects []ProjectConfig
}

func DisplayProjectConfig(p ProjectConfig) {
	v := reflect.ValueOf(p)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: \t%v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

func UnmarshallProjectList(path string) (ProjectList, error) {
	yaml_file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data ProjectList
	err = yaml.Unmarshal(yaml_file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}
