package main

import (
	"github.com/altipla-consulting/configure-dev-machine/internal/installers"
	"github.com/spf13/cobra"
	"libs.altipla.consulting/errors"
)

var filter string

func init() {
	CmdInstall.PersistentFlags().StringVarP(&filter, "filter", "f", "", "Filtra los instaladores a ejecutar")
	CmdRoot.AddCommand(CmdInstall)
}

var CmdInstall = &cobra.Command{
	Use:          "install",
	Short:        "Ejecuta los instaladores",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.Trace(installers.Run(filter))
	},
}
