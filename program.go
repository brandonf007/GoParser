package main

import (
	"fmt"
	"sync"

	appsettings "github.com/brandonf007/GoParser/AppSettings"
	cls "github.com/brandonf007/GoParser/classes"
)

func main() {

	// placeholder variable
	settings := appsettings.AppSettings{}
	// filling the variable with the settings file and env vars
	if err := appsettings.ReadFromFileAndEnv(&settings); err != nil {
		panic(err)
	}

	fmt.Println("--- Begin Parser ---")
	// Process file line by line and store result in two separate hash maps
	if err := cls.ProcessFile(settings.InputFilePath, settings.RegularExpression); err != nil {
		panic(err)
	}
	var wGroup sync.WaitGroup
	// Construct Lists off of the dictionaries
	wGroup.Add(2)
	go cls.ConstructIPList(&wGroup, cls.IPmap)
	go cls.ConstructURLList(&wGroup, cls.URLmap)
	wGroup.Wait()
	// Sort Lists descending by value
	wGroup.Add(2)
	go cls.SortDesc(&wGroup, cls.ILog.IPlist, settings.Output.SortByKeyAdditionally)
	go cls.SortDesc(&wGroup, cls.ULog.URLList, settings.Output.SortByKeyAdditionally)
	wGroup.Wait()
	// Print output to console
	cls.PrintOutputToConsole(cls.ILog, cls.ULog, settings.Output.NumberToDisplay, settings.Output.DisplayUniqueIPs, settings.Output.DisplayUniqueURLs)
	if settings.OutputFile.OutputFilePath != "" {
		if err := cls.PrintOutputToFile(cls.ILog, cls.ULog, settings.Output.NumberToDisplay, settings.Output.DisplayUniqueIPs, settings.Output.DisplayUniqueURLs, settings.OutputFile.OutputFilePath, settings.OutputFile.AppendToOutputFile); err != nil {
			panic(err)
		}
	}

	fmt.Println("--- End Parser ---")
}
