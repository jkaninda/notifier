package cmd

import (
	"jkaninda/notifier/pkg"

	"github.com/spf13/cobra"
)

var sendMessage = &cobra.Command{
	Use:   "sendMessage",
	Short: "Sends Telegram or Mattermost message",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendMessage(cmd)
	},
}

func init() {
	sendMessage.PersistentFlags().StringP("canal", "c", "telegram", "Telegram or Mattermost (telegram, mattermost)")
	sendMessage.PersistentFlags().StringP("message", "m", "", "Message to send")

}
