package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insAZ struct{}

func (ins *insAZ) Name() string {
	return "az"
}

func (ins *insAZ) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insAZ) Install() error {
	script := `
		source /etc/os-release

		sudo apt-get update
		sudo apt-get install -y ca-certificates curl apt-transport-https lsb-release gnupg
		curl -sL https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/microsoft.gpg > /dev/null
		echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ $UBUNTU_CODENAME main" | sudo tee /etc/apt/sources.list.d/azure-cli.list
		sudo apt-get update
		sudo apt-get install -y azure-cli
  `
	if err := run.Shell(script); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (ins *insAZ) BashRC() string {
	return ""
}
