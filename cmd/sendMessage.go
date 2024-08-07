package cmd

import (
	"jkaninda/notifier/pkg"

	"github.com/spf13/cobra"
)

var sendMessage = &cobra.Command{
	Use:   "sendMessage",
	Short: "Sends Telegram notification",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendMessage(cmd)
	},
}

func init() {
	sendEmail.PersistentFlags().StringP("canal", "c", "telegram", "Telegram or Mattermost")
	sendEmail.PersistentFlags().StringP("message", "m", "", "Message to send")

}
