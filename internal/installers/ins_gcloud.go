package installers

import (
	"libs.altipla.consulting/errors"

	"github.com/altipla-consulting/configure-dev-machine/internal/run"
)

type insGcloud struct{}

func (ins *insGcloud) Name() string {
	return "gcloud"
}

func (ins *insGcloud) Check() (*CheckResult, error) {
	return NeedsInstall, nil
}

func (ins *insGcloud) Install() error {
	script := `
    echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | sudo tee /etc/apt/sources.list.d/google-cloud-sdk.list
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
    sudo apt update
    sudo apt install -y google-cloud-sdk kubectl
  `
	if err := run.Shell(script); err != nil {
		return errors.Trace(err)
	}

	if err := run.WriteEnv("CDM_GCLOUD", "1"); err != nil {
		return errors.Trace(err)
	}
	if err := run.WriteEnv("KUBE_EDITOR", "gedit"); err != nil {
		return errors.Trace(err)
	}

	if err := run.WriteAlias("COMPUTE", "compute", "gcloud compute"); err != nil {
		return errors.Trace(err)
	}
	if err := run.WriteAlias("K", "k", "kubectl"); err != nil {
		return errors.Trace(err)
	}
	if err := run.WriteAlias("KLS", "kls", "kubectl config get-contexts"); err != nil {
		return errors.Trace(err)
	}
	if err := run.WriteAlias("KUSE", "kuse", "kubectl config use-context"); err != nil {
		return errors.Trace(err)
	}
	if err := run.WriteAlias("KPODS", "kpods", "kubectl get pods --field-selector=status.phase!=Succeeded -o wide"); err != nil {
		return errors.Trace(err)
	}
	if err := run.WriteAlias("KNODES", "knodes", "kubectl get nodes -o wide"); err != nil {
		return errors.Trace(err)
	}

	script = ` 
      echo "source <(kubectl completion bash | sed 's/kubectl/k/g')" >> ~/.bashrc
    `
	if err := run.Shell(script); err != nil {
		return errors.Trace(err)
	}

	return nil
}
