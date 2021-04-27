package pipeline

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
