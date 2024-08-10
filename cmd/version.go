package cmd

import (
	"fmt"
	"github.com/jkaninda/notifier/pkg"
	"github.com/spf13/cobra"
)

var appVersion = pkg.AppVersion

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		Version()
	},
}

func Version() {
	fmt.Printf("Version: %s \n", appVersion)
	fmt.Println()
}
