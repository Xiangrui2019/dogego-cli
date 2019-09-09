package main

import (
	"github.com/xiangrui2019/dogego-cli/cmd"
	"log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
