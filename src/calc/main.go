package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	// "io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Utc   int64
	Time  string
	Value int64
}

func main() {
	output := flag.String("o", "", "out file")
	input := flag.String("i", "", "in file")

	flag.Parse()
	if *output == "" || *input == "" {
		flag.PrintDefaults()
		fmt.Println("invalid args")
		return
	}

	in, err := os.OpenFile(*input, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println("invalid in file")
		return
	}
	r := csv.NewReader(in)
	r.Comma = '\t'
	r.FieldsPerRecord = 4
	data, err := r.ReadAll()
	zone := time.FixedZone("CST", 8*3600)
	var records []Record
	for _, v := range data[1:] {
		var rec Record
		fmt.Println(v)
		rec.Utc, _ = strconv.ParseInt(v[1], 10, 64)
		t := time.Unix(rec.Utc, 0)
		rec.Time = t.In(zone).String()
		rec.Value, _ = strconv.ParseInt(v[2], 10, 64)
		records = append(records, rec)
	}

	d, _ := json.MarshalIndent(records, "", "")
	ioutil.WriteFile(*output, d, 0777)
	return
}
