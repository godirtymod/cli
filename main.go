package main

import (
	"fmt"

	"bitbucket.org/ai69/amoy"
	w1 "github.com/godirtymod/wrapper1"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
)

var (
	log    *zap.SugaredLogger
	number int
)

func init() {
	displayBuildInfo()
	flag.IntVarP(&number, "number", "n", 1, "Example of number variable")
	flag.Parse()
	log = amoy.SimpleZapSugaredLogger()
}

func main() {
	fmt.Println("Begin1")
	w1.CoreWorkOne()
	fmt.Println("End1")
}
