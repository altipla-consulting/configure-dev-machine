package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(CmdUpdate)
}

var CmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "Imprime el comando de actualización de la herramienta.",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info()
		log.Info("Ejecuta el siguiente comando para actualizar e instalar la última versión:")
		log.Info()
		log.Info("\tcurl https://tools.altipla.consulting/install/configure-dev-machine | bash")
		log.Info()

		return nil
	},
}
