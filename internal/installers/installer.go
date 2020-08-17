package installers

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"libs.altipla.consulting/errors"
)

var (
	register = []Installer{
		new(insApt),
		new(insCI),
		new(insGo),
		new(insNode),
		new(insDC),
		new(insReloader),
		new(insActools),
		new(insGcloud),
		new(insMkcert),
		new(insStern),
	}

	NeedsInstall = &CheckResult{Install: true}
)

type Installer interface {
	Name() string
	Check() (*CheckResult, error)
	Install() error
}

type CheckResult struct {
	Install bool
}

func Run(filter string) error {
	fmt.Println()
	for _, installer := range register {
		if filter != "" && installer.Name() != filter {
			continue
		}

		log.Info("======================================================================")
		log.Info(">>> install ", installer.Name())
		log.Info("======================================================================")

		result, err := installer.Check()
		if err != nil {
			return errors.Trace(err)
		}
		if result != nil {
			if result.Install {
				if err := installer.Install(); err != nil {
					return errors.Trace(err)
				}
			}
		}

		fmt.Println()
		fmt.Println()
	}

	log.Info("======================================================================")
	log.Info("======================================================================")
	log.Info()
	log.Info("Finished successfully!")
	log.Info()

	return nil
}
