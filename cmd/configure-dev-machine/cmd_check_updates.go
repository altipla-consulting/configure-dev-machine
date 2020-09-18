package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"libs.altipla.consulting/errors"
)

func init() {
	CmdRoot.AddCommand(CmdCheckUpdates)
}

var CmdCheckUpdates = &cobra.Command{
	Use:   "check-updates",
	Short: "Comprueba si hay actualizaciones de esta herramienta.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Jenkins y el entorno de desarrollo no deben comprobar la versión
		if Version == "dev" || os.Getenv("CI") == "true" {
			return nil
		}

		configdir, err := os.UserConfigDir()
		if err != nil {
			return errors.Trace(err)
		}
		filename := filepath.Join(configdir, "configure-dev-machine", "last-update-check.txt")

		if err := os.MkdirAll(filepath.Dir(filename), 0700); err != nil {
			return errors.Trace(err)
		}

		lastUpdate := time.Time{}
		if content, err := ioutil.ReadFile(filename); err != nil && !os.IsNotExist(err) {
			return errors.Trace(err)
		} else if err == nil {
			if err := lastUpdate.UnmarshalText(content); err != nil {
				return errors.Trace(err)
			}
		}

		if time.Now().Sub(lastUpdate) > 1*time.Hour {
			reply, err := http.Get("https://tools.altipla.consulting/version-manifest/configure-dev-machine")
			if err != nil {
				return errors.Trace(err)
			}
			defer reply.Body.Close()
			if reply.StatusCode != http.StatusOK {
				return errors.Errorf("unexpected status: %s", reply.Status)
			}
			body, err := ioutil.ReadAll(reply.Body)
			if err != nil {
				return errors.Trace(err)
			}
			expected := strings.TrimSpace(string(body))

			if expected != Version {
				log.WithFields(log.Fields{"current": Version, "latest": expected}).Error("configure-dev-machine has been updated")

				log.Warning()
				log.Warning("Run the following command to install the latest version:")
				log.Warning()
				log.Warning("\tcurl https://tools.altipla.consulting/install/configure-dev-machine | bash")
				log.Warning()

				return nil
			}

			check, err := time.Now().MarshalText()
			if err != nil {
				return errors.Trace(err)
			}
			if err := ioutil.WriteFile(filename, check, 0600); err != nil {
				return errors.Trace(err)
			}
		}

		return nil
	},
}
