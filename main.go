package main

import (
	"flag"
	"github.com/mm-ooto/awsEmailSmsSend/config"
)

var (
	runMode string
)

func init() {
	flag.StringVar(&runMode, "runMode", "dev", "managed environment")
	flag.Parse()
}

func main() {
	config.LoadConfig(runMode)

}
