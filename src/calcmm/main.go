package main

import (
	"encoding/json"
	"flag"
	"fmt"
	hum "github.com/dustin/go-humanize"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	args := flag.Args()
	logPath := args[0]
	rets := [][]Record{}

	filepath.Walk(logPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			records, err1 := readRecords(path)
			if err1 != nil {
				return err
			}
			rets = append(rets, records)
		}
		fmt.Println(path)
		return err
	})

	result := merge(rets)

	fmt.Println("length", len(result))
	pos := int(float32(len(result)) * 0.95)

	var rec MergeRecord
	fmt.Println("-----------avg-------------")
	rec = result[pos]
	fmt.Println("95%", pos, rec.Time, hum.Comma(rec.Total), hum.Bytes(uint64(rec.Total)))
	fmt.Printf("%#v\n", rec.PerLog)
	pos = len(result) - 4
	rec = result[pos]
	fmt.Println("4th", pos, rec.Time, hum.Comma(rec.Total), hum.Bytes(uint64(rec.Total)))
	fmt.Printf("%#v\n", rec.PerLog)
	pos = len(result) - 1
	rec = result[pos]
	fmt.Println("max", pos, rec.Time, hum.Comma(rec.Total), hum.Bytes(uint64(rec.Total)))
	fmt.Printf("%#v\n", rec.PerLog)
	data, _ := json.MarshalIndent(result, "", "")
	ioutil.WriteFile("all.log", data, 0600)
	return
}
