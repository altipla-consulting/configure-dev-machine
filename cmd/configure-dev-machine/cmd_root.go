package main

import (
	"github.com/spf13/cobra"
)

var CmdRoot = &cobra.Command{
	Use:          "configure-dev-machine",
	Short:        "Instalador y configurador para los ordenadores de desarrolladores",
	SilenceUsage: true,
}
