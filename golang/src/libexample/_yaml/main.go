package main

import (
	"fmt"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var data = `
test:
  pre:
    - go get github.com/axw/gocov/gocov
    - go get github.com/mattn/goveralls
    - mkdir -p $HOME/.go_workspace/src/_/home/ubuntu/
    - ln -s $HOME/$CIRCLE_PROJECT_REPONAME $HOME/.go_workspace/src/_/home/ubuntu/
  override:
    - go test -v ./docomo
    - goveralls -v -service=circle-ci -repotoken $COVERALLS_TOKEN ./docomo
`

type CircleCIConfig struct {
	Test struct{
		Pre []string
		Override []string
	}
}

func main() {
	t := CircleCIConfig{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))


	for _, p := range t.Test.Pre {
		fmt.Println(p)
	}
}
