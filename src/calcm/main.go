package main

import (
	"encoding/json"
	"flag"
	"fmt"
	hum "github.com/dustin/go-humanize"
	"io/ioutil"
)

func main() {
	output := flag.String("o", "", "out file")

	flag.Parse()
	if *output == "" {
		flag.PrintDefaults()
		fmt.Println("invalid args")
		return
	}
	args := flag.Args()
	rets := [][]RecordHour{}
	for _, v := range args {
		ret, total, err := Records(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("amount", v, hum.Comma(total))
		rets = append(rets, ret)
	}

	result := merge(rets)

	fmt.Println("length", len(result))
	pos := int(float32(len(result)) * 0.95)
	amount := calcThoughput(result)
	fmt.Println("total amount", hum.Comma(amount), hum.Bytes(uint64(amount)))
	var rec *MergeHour
	fmt.Println("-----------avg-------------")
	rec = result[pos]
	fmt.Println(pos, rec, hum.Comma(rec.Total), hum.Bytes(uint64(rec.Total)))
	rec = result[len(result)-1]
	fmt.Println(pos, rec, hum.Comma(rec.Total), hum.Bytes(uint64(rec.Total)))
	d, _ := json.MarshalIndent(result, "", "")
	ioutil.WriteFile(*output, d, 0777)
	return
}
