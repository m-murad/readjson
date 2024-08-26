package main

import (
	"log"
	"os"
	"readjson/pkg/readjson"
)

func main() {
	app := readjson.NewApp(&log.Logger{}, "/some/directory", os.Stdout)
	// we can use something other than `os.Stdout` to record results
	app.Start()
}
