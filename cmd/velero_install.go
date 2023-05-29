package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installVelero = &cobra.Command {
	Use: "install",
	Short: "Install Velero in the cluster",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		
	}
}