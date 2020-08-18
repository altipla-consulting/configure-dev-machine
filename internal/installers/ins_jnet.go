package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insJnet struct{}

func (ins *insJnet) Name() string {
	return "jnet"
}

func (ins *insJnet) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insJnet) Install() error {
	script := `
    mkdir -p ~/bin
    curl -L -o ~/bin/jnet $(curl --silent 'https://api.github.com/repos/altipla-consulting/jnet/releases/latest' | jq -r '.assets[0].browser_download_url')
    chmod +x ~/bin/jnet
  `
	return errors.Trace(run.Shell(script))
}
