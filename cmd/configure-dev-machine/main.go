package main

import (
	log "github.com/sirupsen/logrus"
	"libs.altipla.consulting/errors"
)

var Version = "dev"

func main() {
	if err := CmdRoot.Execute(); err != nil {
		log.Fatal(errors.Stack(err))
	}
}
