package installers

import (
	"os"

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

	if os.Getenv("CDM_GCLOUD") == "" {
		if err := r.WriteEnv("CDM_GCLOUD", "1"); err != nil {
			return errors.Trace(err)
		}
		if err := r.WriteEnv("KUBE_EDITOR", "gedit"); err != nil {
			return errors.Trace(err)
		}

		if err := r.WriteAlias("compute", "gcloud compute"); err != nil {
			return errors.Trace(err)
		}
		if err := r.WriteAlias("k", "kubectl"); err != nil {
			return errors.Trace(err)
		}
		if err := r.WriteAlias("kls", "kubectl config get-contexts"); err != nil {
			return errors.Trace(err)
		}
		if err := r.WriteAlias("kuse", "kubectl config use-context"); err != nil {
			return errors.Trace(err)
		}
		if err := r.WriteAlias("kpods", "kubectl get pods --field-selector=status.phase!=Succeeded -o wide"); err != nil {
			return errors.Trace(err)
		}
		if err := r.WriteAlias("knodes", "kubectl get nodes -o wide"); err != nil {
			return errors.Trace(err)
		}

		script := ` 
      echo "source <(kubectl completion bash | sed 's/kubectl/k/g')" >> ~/.bashrc
    `
		if err := run.Shell(script); err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}
