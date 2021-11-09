package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/signaux-faibles/libinpirncs"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(4)
	}
	reference := os.Args[1]
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	bilan, err := libinpirncs.ParseBilan(data, reference)
	if err != nil {
		os.Exit(2)
	}
	output, err := json.MarshalIndent(bilan, " ", " ")
	if err != nil {
		os.Exit(3)
		return
	}
	fmt.Printf("%s", output)
}
