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

type RecordHour struct {
	Utc      int64
	Time     string
	ValueMin int64
	ValueMax int64
	ValueAvg int64
}

var zone *time.Location = time.FixedZone("CST", 8*3600)

func valid(utc int64) bool {
	t := time.Unix(utc, 0)
	start := time.Date(2014, 4, 1, 0, 0, 0, 0, zone)
	end := time.Date(2014, 5, 1, 0, 0, 0, 1, zone)

	return t.After(start) && t.Before(end)
}

func Records(path string) (records []RecordHour, err error) {
	in, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer in.Close()
	r := csv.NewReader(in)
	r.Comma = '\t'
	r.FieldsPerRecord = 6
	data, err := r.ReadAll()

	for _, v := range data[1:] {
		var rec RecordHour
		rec.Utc, _ = strconv.ParseInt(v[1], 10, 64)
		if !valid(rec.Utc) {
			continue
		}
		t := time.Unix(rec.Utc, 0)
		rec.Time = t.In(zone).String()
		rec.ValueMin, _ = strconv.ParseInt(v[3], 10, 64)
		rec.ValueAvg, _ = strconv.ParseInt(v[4], 10, 64)
		rec.ValueMax, _ = strconv.ParseInt(v[5], 10, 64)
		records = append(records, rec)
	}

	return
}
