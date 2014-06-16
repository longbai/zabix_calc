package main

import (
	// "fmt"
	"sort"
)

type MergeRecord struct {
	Utc    int64
	Time   string
	Total  int64
	PerLog []int64
}

type Array []MergeRecord

func (p Array) Len() int           { return len(p) }
func (p Array) Less(i, j int) bool { return p[i].Total < p[j].Total }
func (p Array) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Array) Sort()              { sort.Sort(p) }

func merge(args [][]Record) (ret []MergeRecord) {
	ret = make([]MergeRecord, pointNumber)
	first := args[0]
	length := len(args)
	for i, v := range first {
		ret[i].Utc = v.Utc
		ret[i].Time = v.printTime()
		ret[i].Total = v.Value
		ret[i].PerLog = []int64{v.Value}
		for j := 1; j < length; j++ {
			val := args[j][i].Value
			ret[i].Total += val
			ret[i].PerLog = append(ret[i].PerLog, val)
		}
	}
	Array(ret).Sort()
	return
}
