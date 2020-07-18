package main

import (
	log "github.com/sirupsen/logrus"
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/installers"
)

func main() {
	if err := installers.Run(); err != nil {
		log.Fatal(errors.Stack(err))
	}
}
