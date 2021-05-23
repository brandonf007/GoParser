package classes

import (
	"sort"
	"sync"
)

type KeyVal struct {
	Key string
	Val int
}

type ProcessedIPLog struct {
	IPlist        []KeyVal
	UniqueIPCount int
}

type ProcessedURLLog struct {
	URLList        []KeyVal
	UniqueURLCount int
}

var ILog ProcessedIPLog
var ULog ProcessedURLLog

// Threaded task to build a an array of KeyVal pair containing the IP addresses and the number of hits
func ConstructIPList(wGroup *sync.WaitGroup, dict map[string]int) {
	defer wGroup.Done()
	ILog.UniqueIPCount = len(dict)
	for key, value := range dict {
		ILog.IPlist = append(ILog.IPlist, KeyVal{key, value})
	}
}

// Threaded task to build a an array of KeyVal pair containing the URLs and the number of occurrences
func ConstructURLList(wGroup *sync.WaitGroup, dict map[string]int) {
	defer wGroup.Done()
	ULog.UniqueURLCount = len(dict)
	for key, value := range dict {
		ULog.URLList = append(ULog.URLList, KeyVal{key, value})
	}
}

// This will sort the Array of keyval pair by the value in descending order, if sortAndThen is provided it will also sort by the key
func SortDesc(wGroup *sync.WaitGroup, list []KeyVal, sortAndThen bool) {
	defer wGroup.Done()
	sort.Slice(list, func(i, j int) bool {
		if sortAndThen {
			if list[i].Val == list[j].Val {
				return list[i].Key > list[j].Key
			} else {
				return list[i].Val > list[j].Val
			}
		} else {
			return list[i].Val > list[j].Val
		}
	})
}
