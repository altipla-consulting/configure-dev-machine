package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insReloader struct{}

func (ins *insReloader) Name() string {
	return "reloader"
}

func (ins *insReloader) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insReloader) Install() error {
	script := `
    curl -L -o ~/bin/reloader $(curl --silent 'https://api.github.com/repos/altipla-consulting/reloader/releases/latest' | jq -r '.assets[0].browser_download_url')
    chmod +x ~/bin/reloader
  `
	return errors.Trace(run.Shell(script))
}
