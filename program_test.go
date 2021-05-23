package main

import (
	"sync"

	cls "github.com/brandonf007/GoParser/classes"
)

//This runs on a log file provided to me
func Example_one() {
	// Process file line by line and store result in two separate hash maps
	if err := cls.ProcessFile(".\\TestFiles\\ProvidedExample.log", "(^(?:[0-9]{1,3}\\.){3}[0-9]{1,3})(?:.+)(?:\"{1}(?:GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE){1}\\s{1}(.+)\\s{1}(?:HTTP/1.1|HTTP/1.0|HTTP/2|HTTP/3){1}\"{1})(?:.+$)"); err != nil {
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
	go cls.SortDesc(&wGroup, cls.ILog.IPlist, true)
	go cls.SortDesc(&wGroup, cls.ULog.URLList, true)
	wGroup.Wait()
	// Print output to console
	cls.PrintOutputToConsole(cls.ILog, cls.ULog, 3, true, false)
	//Output: Number of unique IP Addresses:11
	//Most Active IP Addresses:168.41.191.40 - HITS:4
	//Most Active IP Addresses:72.44.32.10 - HITS:3
	//Most Active IP Addresses:50.112.00.11 - HITS:3
	//Most visited URLS:/docs/manage-websites/ - Occurences:2
	//Most visited URLS:http://example.net/faq/ - Occurences:1
	//Most visited URLS:http://example.net/blog/category/meta/ - Occurences:1

}

// This runs on a log file I have found
func Example_two() {
	// Process file line by line and store result in two separate hash maps
	// clear dictionaries so that they can be repopulated
	cls.IPmap = make(map[string]int)
	cls.URLmap = make(map[string]int)
	if err := cls.ProcessFile(".\\TestFiles\\ISLogFile.log", "(^(?:[0-9]{1,3}\\.){3}[0-9]{1,3})(?:.+)(?:\"{1}(?:GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE){1}\\s{1}(.+)\\s{1}(?:HTTP/1.1|HTTP/1.0|HTTP/2|HTTP/3){1}\"{1})(?:.+$)"); err != nil {
		panic(err)
	}
	var wGroup sync.WaitGroup
	// clear lists so that they can be repopulated
	cls.ILog = cls.ProcessedIPLog{}
	cls.ULog = cls.ProcessedURLLog{}
	// Construct Lists off of the dictionaries
	wGroup.Add(2)
	go cls.ConstructIPList(&wGroup, cls.IPmap)
	go cls.ConstructURLList(&wGroup, cls.URLmap)
	wGroup.Wait()
	wGroup.Add(2)
	go cls.SortDesc(&wGroup, cls.ILog.IPlist, true)
	go cls.SortDesc(&wGroup, cls.ULog.URLList, true)
	wGroup.Wait()
	// Print output to console
	cls.PrintOutputToConsole(cls.ILog, cls.ULog, 3, true, false)
	//Output: Number of unique IP Addresses:1753
	//Most Active IP Addresses:66.249.73.135 - HITS:482
	//Most Active IP Addresses:46.105.14.53 - HITS:364
	//Most Active IP Addresses:130.237.218.86 - HITS:357
	//Most visited URLS:/favicon.ico - Occurences:807
	//Most visited URLS:/style2.css - Occurences:546
	//Most visited URLS:/reset.css - Occurences:538

}
