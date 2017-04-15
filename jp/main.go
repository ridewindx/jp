package main

import (
	"os"
	"encoding/json"
	"fmt"
	"bytes"
	"flag"
	"github.com/ridewindx/jp"
)

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage of jp:
	jp '{.jsonpath}' < filename.json
	cat filename.json | jp '{.jsonpath}'`)
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var input interface{}
	err := json.NewDecoder(os.Stdin).Decode(&input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	j := jsonpath.New("")
	err = j.Parse(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	buf := new(bytes.Buffer)
	err = j.Execute(buf, input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(buf.String())
	return
}
