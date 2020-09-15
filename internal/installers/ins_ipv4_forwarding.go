package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insIPV4Forwarding struct{}

func (ins *insIPV4Forwarding) Name() string {
	return "ipv4-forwarding"
}

func (ins *insIPV4Forwarding) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insIPV4Forwarding) Install() error {
	script := `
		sudo sysctl -w net.ipv4.ip_forward=1
		sudo sysctl --system
  `
	return errors.Trace(run.Shell(script))
}
