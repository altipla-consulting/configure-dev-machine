package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insCI struct{}

func (ins *insCI) Name() string {
	return "ci"
}

func (ins *insCI) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insCI) Install() error {
	script := `
    mkdir -p ~/bin
    curl -L -o ~/bin/ci $(curl --silent 'https://api.github.com/repos/altipla-consulting/ci/releases/latest' | jq -r '.assets[0].browser_download_url')
    chmod +x ~/bin/ci
  `
	return errors.Trace(run.Shell(script))
}

func (ins *insCI) BashRC() string {
	return ""
}
