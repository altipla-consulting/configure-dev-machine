package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insApt struct{}

func (ins *insApt) Name() string {
	return "apt"
}

func (ins *insApt) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insApt) Install() error {
	script := `
    sudo apt update
    sudo apt install -y wget tar curl autoconf jq git build-essential
  `
	return errors.Trace(run.Shell(script))
}
