package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type AuthorizationPolicy struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
		Labels    struct {
			TargetService string `yaml:"targetService"`
		} `yaml:"labels"`
		Annotations struct {
			RouterDeisIoDomains string `yaml:"router.fronteira.io/domains"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec Spec `yaml:"spec"`
}

type Spec struct {
	Operation []Operation `yaml:"operations"`
	Target    string      `yaml:"target"`
}
type Operation struct {
	Method string   `yaml:"method"`
	Path   string   `yaml:"path"`
	Roles  []string `yaml:"roles"`
}

func NewAuthPolicy(path string) *AuthorizationPolicy {
	authPolicy := AuthorizationPolicy{}
	file := authPolicy.fileReader(path)
	authPolicy.YAMLUnmarshall(file)
	return &authPolicy
}
func (a *AuthorizationPolicy) YAMLUnmarshall(file []byte) {
	err := yaml.Unmarshal(file, &a)
	if err != nil {
		panic(err)
	}
}

func (a *AuthorizationPolicy) fileReader(path string) []byte {
	filename, _ := filepath.Abs(path)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	return yamlFile
}
