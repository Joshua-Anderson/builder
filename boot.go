package main

import (
	"log"
	"os"
	"runtime"

	"github.com/deis/builder/fetcher"
	"github.com/deis/builder/pkg"
	"github.com/kelseyhightower/envconfig"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var conf Config
	if err := envconfig.Process("builder", &conf); err != nil {
		log.Fatalf("error fetching config [%s]", err)
		os.Exit(1)
	}
	log.Printf("starting fetcher on port %d", conf.FetcherPort)
	go fetcher.Serve(conf.FetcherPort)
	os.Exit(pkg.Run(conf.SSHHostIP, conf.SSHHostPort, "boot"))
}
