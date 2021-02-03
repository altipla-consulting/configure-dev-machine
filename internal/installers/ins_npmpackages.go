package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insNpmpackages struct{}

func (ins *insNpmpackages) Name() string {
	return "npmpackages"
}

func (ins *insNpmpackages) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insNpmpackages) Install() error {
	script := `
	sudo npm install -g npm@latest
	sudo npm install -g yarn@latest
	sudo npm install -g netlify-cli@latest
  `
	return errors.Trace(run.Shell(script))
}
