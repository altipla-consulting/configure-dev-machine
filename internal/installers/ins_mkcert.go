package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insMkcert struct{}

func (ins *insMkcert) Name() string {
	return "mkcert"
}

func (ins *insMkcert) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insMkcert) Install() error {
	script := `
    curl -L -o ~/bin/mkcert https://github.com/FiloSottile/mkcert/releases/download/v1.4.1/mkcert-v1.4.1-linux-amd64
    chmod +x ~/bin/mkcert
    mkcert -install
  `
	return errors.Trace(run.Shell(script))
}
