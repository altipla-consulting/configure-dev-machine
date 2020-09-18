package installers

import (
	"os"

	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insUpdater struct{}

func (ins *insUpdater) Name() string {
	return "updater"
}

func (ins *insUpdater) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insUpdater) Install() error {
	if os.Getenv("CONFIGURE_DEV_MACHINE_UPDATER") == "" {
		script := `
      echo 'export CONFIGURE_DEV_MACHINE_UPDATER=true' >> ~/.bashrc
			echo 'configure-dev-machine check-updates' >> ~/.bashrc
    `
		if err := run.Shell(script); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
