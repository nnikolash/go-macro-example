package main

import (
	"encoding/json"
	"fmt"
)

//go:generate bin/tools/include macros/logging.h

var LoggingEnabled = true

func main() {
	fmt.Printf("Program started\n")
	defer fmt.Printf("Program stopped\n")

	LOG("Hello, World! --- %v", toJSON(BigStructure{}))
}

type BigStructure struct {
	Field1 [30]int
}

func toJSON(v any) string {
	j, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(j)
}
