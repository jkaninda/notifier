package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "notifier [Command]",
	Short:   "Send notification using Email or Telegram",
	Long:    `Send notification using Email or Telegram`,
	Example: "",
	Version: appVersion,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(sendMessage)
	rootCmd.AddCommand(sendMail)
	rootCmd.PersistentFlags().StringP("env-file", "", "", "Environment variable file")

}
