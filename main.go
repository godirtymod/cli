package main

import (
	"fmt"
	"os"
	"time"

	"bitbucket.org/ai69/amoy"
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
	pwd, _ := os.Getwd()
	host, _ := os.Hostname()
	log.Infow("Hello New World", "hostname", host, "work_dir", pwd, "num", number)
	fmt.Printf("🌋: Hello, World!\n⏰: %s\n", time.Now().Format("2006-01-02T15:04:05-0700"))
}
