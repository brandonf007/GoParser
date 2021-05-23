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

func addMap(dict map[string]int, key string) {
	if j, exists := dict[key]; exists {
		dict[key] = j + 1
	} else {
		dict[key] = 1
	}
}

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
	for scanner.Scan() {
		var match = reg.FindStringSubmatch(scanner.Text())
		addMap(IPmap, match[1])
		addMap(URLmap, match[2])
		Successful++
	}
	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "Failed to process file")
	}
	return err
}
