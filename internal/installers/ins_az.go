package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insAZ struct{}

func (ins *insAZ) Name() string {
	return "az"
}

func (ins *insAZ) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insAZ) Install() error {
	script := `
    curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash
  `
	if err := run.Shell(script); err != nil {
		return errors.Trace(err)
	}

	return nil
}
