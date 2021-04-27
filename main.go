package main

import (
	"fmt"
	"github.com/mrcyna/pipeline-to-bson/pipeline"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	content := strings.TrimSpace(string(data))
	if !pipeline.Validate(content) {
		fmt.Println("File is invalid", err)
		return
	}

	fmt.Println("OK!")
}
