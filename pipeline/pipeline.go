package pipeline

import (
	"regexp"
	"strings"
)

// Validate will check the input content for valid pipeline statement
func Validate(content string) bool {

	// Pipeline should not be empty
	if len(content) == 0 {
		return false
	}

	// Pipeline should not start with any char other than "["
	if string(content[0]) != "[" {
		return false
	}

	// Pipeline should not end with any char other than "]"
	if string(content[len(content) - 1]) != "]" {
		return false
	}

	// Pipeline building statements should be opened & closed correctly
	count := make(map[string]int, 2)
	count["square"] = 0
	count["mustache"] = 0

	for _, c := range content {
		ch := string(c)

		if ch == "[" {
			count["square"]++
		}

		if ch == "]" {
			count["square"]--
		}

		if ch == "{" {
			count["mustache"]++
		}

		if ch == "}" {
			count["mustache"]--
		}

		// Open and Close should not be less than zero
		if count["square"] < 0 || count["mustache"] < 0 {
			return false
		}
	}

	// At the end we should have no extra open or closed bracket
	if count["square"] != 0 || count["mustache"] != 0 {
		return false
	}

	return true
}

// Transform will convert JSON like pipeline into Golang bson type
func Transform(content string) string {
	content = strings.ReplaceAll(content, "{", "bson.M{")
	content = strings.ReplaceAll(content, "[", "bson.A{")
	content = strings.ReplaceAll(content, "]", "}")

	re := regexp.MustCompile(`bson.M{\s*(\$[^:]*):`)
	content = re.ReplaceAllString(content, "bson.M{\"$1\":")

	output := ""
	for _, line := range strings.Split(content,"\n") {
		statement := strings.TrimSpace(line)
		if statement[len(statement)-1:] == "}" {
			output += statement + ",\n"
		} else {
			output += statement + "\n"
		}
	}

	return output
}
