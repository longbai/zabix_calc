package main

import (
	"encoding/json"
	"flag"
	"fmt"
	// "io"
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
		fmt.Println(v)
		ret, err := Records(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		rets = append(rets, ret)
	}

	result := merge(rets)

	fmt.Println("length", len(result))
	pos := int(float32(len(result)) * 0.95)

	fmt.Println("-----------avg-------------")
	fmt.Println(pos, result[pos])
	fmt.Println(pos+1, result[pos+1])
	fmt.Println(len(result)-1, result[len(result)-1])
	d, _ := json.MarshalIndent(result, "", "")
	ioutil.WriteFile(*output, d, 0777)

	fmt.Println("-----------max-------------")
	result = mergeMax(rets)
	fmt.Println(pos, result[pos])
	fmt.Println(pos+1, result[pos+1])
	fmt.Println(len(result)-1, result[len(result)-1])

	return
}
