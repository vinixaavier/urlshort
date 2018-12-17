package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// PathToURL struct to parse content of the yaml file.
type PathToURL struct {
	Path string `yaml: "path"`
	URL  string `yaml: "url"`
}

// PathsToURL struct means a PathToURL list
type PathsToURL []PathToURL

// New initialize a empty struct
func New() *PathsToURL {
	return &PathsToURL{}
}

// YAML function receive a yaml buffer and redirect incoming request context to URL specify.
func (paths *PathsToURL) YAML(file *string, fallback http.Handler) (http.HandlerFunc, error) {

	reader, _ := os.Open(*file)
	defer reader.Close()
	buffer, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(buffer, &paths)

	if err != nil {
		log.Fatal(err)
	}

	mapper := make(map[string]string)
	for _, value := range *paths {
		mapper[value.Path] = value.URL
	}

	return Map(mapper, fallback), nil

}
