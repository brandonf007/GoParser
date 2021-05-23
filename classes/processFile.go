package classes

import (
	"bufio"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

var IPmap = make(map[string]int)
var URLmap = make(map[string]int)
var Successful int

// Add key to dictionary if there is a clash increase the count associated to the key
func addMap(dict map[string]int, key string) {
	if j, exists := dict[key]; exists {
		dict[key] = j + 1
	} else {
		dict[key] = 1
	}
}

// This will process a file passed to it line by line, matching the line to the provided regular expression
func ProcessFile(fileInput string, regexPattern string) error {
	file, err := os.Open(fileInput)
	if err != nil {
		return errors.Wrap(err, "Failed to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile(regexPattern)
	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "Failed to compile regex")
	}
	// Read the file line for line
	for scanner.Scan() {
		var match = reg.FindStringSubmatch(scanner.Text())
		if len(match) != 3 {
			return errors.Wrap(errors.New("Regular Expression - error"), "Regular Expression should contain two groups the first identifying the IP Address, and the second identifying the URL")
		}
		addMap(IPmap, match[1])
		addMap(URLmap, match[2])
		Successful++
	}
	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "Failed to process file")
	}
	return err
}
