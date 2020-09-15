package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insEnospc struct{}

func (ins *insEnospc) Name() string {
	return "enospc"
}

func (ins *insEnospc) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insEnospc) Install() error {
	script := `
		sudo sysctl -w fs.inotify.max_user_watches=524288
		sudo sysctl --system
  `
	return errors.Trace(run.Shell(script))
}
