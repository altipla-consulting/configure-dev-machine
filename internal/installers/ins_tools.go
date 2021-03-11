package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insTools struct{}

func (ins *insTools) Name() string {
	return "reloader"
}

func (ins *insTools) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insTools) Install() error {
	script := `
		curl https://tools.altipla.consulting/install/tools | bash
  `
	return errors.Trace(run.Shell(script))
}

func (ins *insTools) BashRC() string {
	return ""
}
