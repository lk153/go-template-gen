package yaml

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"

	envpkg "github.com/lk153/go-template-gen/internal/env"
)

type (
	FileProcessor interface {
		Read(filename string)
	}

	YamlProcessor struct {
		data any
	}
)

func InitYamlProcessor(data envpkg.EnvVars) YamlProcessor {
	return YamlProcessor{
		data: data,
	}
}

func (y YamlProcessor) Read(filename string) {
	// Read config file into byte slice
	var finalData bytes.Buffer
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	t := template.New("config")
	t, err = t.Parse(string(fileData))
	if err != nil {
		panic(err)
	}

	t.Execute(&finalData, y.data)
	fmt.Println(finalData.String())
}
