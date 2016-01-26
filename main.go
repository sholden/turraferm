package main

import (
	"os"
	"fmt"
	"encoding/json"
	"github.com/hashicorp/terraform/config"
)

func main() {
	os.Exit(realMain())
}

//var json_data = make(map[string]interface{})

func jsonify(path string) int {
	if len(path) < 1 {
		panic("Need a path, dummy")
	}

	tfconfig, err := config.LoadDir(path)
	if err != nil {
		panic(err)
	}

	cfgjson, err := json.MarshalIndent(tfconfig, "", "    ")
	fmt.Println(string(cfgjson))

	return 0
}

func realMain() int {
	args := os.Args[1:]

	if args[0] == "jsonify" {
		return jsonify(args[1])
	}
	return 1
}