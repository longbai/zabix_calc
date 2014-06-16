package main

import (
	"encoding/csv"
	// "encoding/json"
	// "flag"
	// "fmt"
	// "io"
	// "io/ioutil"
	"os"
	"strconv"
	"time"
)

var zone *time.Location = time.FixedZone("CST", 8*3600)
var start time.Time = time.Date(2014, 5, 1, 0, 0, 0, 0, zone)
var startUnix int64 = start.Unix()
var end time.Time = time.Date(2014, 6, 1, 0, 0, 0, 1, zone)
var endUnix int64 = end.Unix()

const interval = 5 * 60

const pointNumber = 288 * 31

type Record struct {
	Utc   int64
	Value int64
}

func valid(utc int64) bool {
	return utc > startUnix && utc < endUnix
}

func init5MinRecords() []Record {
	recs := make([]Record, pointNumber)
	fromUtc := startUnix + interval
	for i := 0; i < pointNumber; i++ {
		recs[i].Utc = fromUtc + int64(i*interval)
	}
	return recs
}

func (l *Record) printTime() string {
	t := time.Unix(l.Utc, 0)
	return t.In(zone).String()
}

func readMinuteRecords(path string) (records []Record, err error) {
	in, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer in.Close()
	r := csv.NewReader(in)
	r.Comma = '\t'
	r.FieldsPerRecord = 4
	data, err := r.ReadAll()
	for _, v := range data[1:] {
		var rec Record
		rec.Utc, _ = strconv.ParseInt(v[1], 10, 64)
		if !valid(rec.Utc) {
			continue
		}
		rec.Value, _ = strconv.ParseInt(v[2], 10, 64)
		records = append(records, rec)
	}
	return
}

func readRecords(path string) (recordsOf5 []Record, err error) {
	minRecord, err := readMinuteRecords(path)
	if err != nil {
		return
	}
	var tempCount int64
	var temp5MinValues int64
	var curPos int
	recordsOf5 = init5MinRecords()
	for _, v := range minRecord {
		var pos int = int((v.Utc - startUnix) / interval)
		if curPos == 0 {
			curPos = pos
		}
		if pos != curPos {
			if temp5MinValues != 0 {
				recordsOf5[curPos].Value = temp5MinValues / tempCount
			}
			curPos = pos
			tempCount = 0
			temp5MinValues = 0
		} else {
			tempCount++
			temp5MinValues += v.Value
		}
	}

	return
}
