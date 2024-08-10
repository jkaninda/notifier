package cmd

import (
	"jkaninda/notifier/pkg"

	"github.com/spf13/cobra"
)

var sendMail = &cobra.Command{
	Use:   "sendMail",
	Short: "Sends email",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendEmail(cmd)
	},
}

func init() {
	sendMail.PersistentFlags().StringP("to", "t", "", "Destination email")
	sendMail.PersistentFlags().StringP("subject", "", "", "Email subject")
	sendMail.PersistentFlags().StringP("message", "", "", "Email message")
	sendMail.PersistentFlags().BoolP("insecure", "", false, "Disable StartTLS")
	sendMail.PersistentFlags().StringP("attach", "", "", "Attach a file")
}
