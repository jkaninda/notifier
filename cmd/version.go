package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"jkaninda/notifier/pkg"
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
