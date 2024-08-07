package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var appVersion = os.Getenv("VERSION")

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
