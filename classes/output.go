package classes

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func PrintOutputToFile(iLog ProcessedIPLog, uLog ProcessedURLLog, numberOfHits int, displayUniqueIp bool, displayUniqueURL bool, outputFile string, append bool) error {
	var file *os.File
	var err error
	if !append {
		file, err = os.Create(outputFile)
		if err != nil {
			file.Close()
			return errors.Wrap(err, fmt.Sprintf("Failed to create file: %v", outputFile))
		}
	} else {
		file, err = os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			file.Close()
			file, err = os.Create(outputFile)
			if err != nil {
				file.Close()
				return errors.Wrap(err, fmt.Sprintf("Failed to append to file: %v", outputFile))
			}
		}
	}

	l, err := file.WriteString(generateOutput(iLog, uLog, numberOfHits, displayUniqueIp, displayUniqueURL))
	if err != nil {
		file.Close()
		return errors.Wrap(err, fmt.Sprintf("Failed to write to file: %v", outputFile))
	}
	if l <= 0{
		return errors.Wrap(err, fmt.Sprintf("No Bytes written to file: %v", outputFile))
	}
	file.Close()
	return err
}

func PrintOutputToConsole(iLog ProcessedIPLog, uLog ProcessedURLLog, numberOfHits int, displayUniqueIp bool, displayUniqueURL bool) {
	fmt.Println(generateOutput(iLog, uLog, numberOfHits, displayUniqueIp, displayUniqueURL))
}

func generateOutput(iLog ProcessedIPLog, uLog ProcessedURLLog, numberOfHits int, displayUniqueIp bool, displayUniqueURL bool) string {
	output := ""
	counter := 0
	if displayUniqueIp {
		output += fmt.Sprintf("Number of unique IP Addresses:%v\n", iLog.UniqueIPCount)
	}
	for _, i := range iLog.IPlist {
		counter++
		output += fmt.Sprintf("Most Active IP Addresses:%v - HITS:%v\n", i.Key, i.Val)
		if counter == numberOfHits {
			break
		}
	}
	counter = 0
	if displayUniqueURL {
		output += fmt.Sprintf("Number of unique URLs:%v\n", uLog.UniqueURLCount)
	}
	for _, i := range uLog.URLList {
		counter++
		output += fmt.Sprintf("Most visited URLS:%v - Occurences:%v\n", i.Key, i.Val)
		if counter == numberOfHits {
			break
		}
	}
	output += "\n"
	return output
}
