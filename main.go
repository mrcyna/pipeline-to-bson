package main

import (
	"bufio"
	"fmt"
	"github.com/mrcyna/pipeline-to-bson/pipeline"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the filepath:\n")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	content := strings.TrimSpace(string(data))
	if !pipeline.Validate(content) {
		fmt.Println("File is invalid", err)
		return
	}

	fmt.Println(pipeline.Transform(content))
}
