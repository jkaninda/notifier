package cmd

import (
	"jkaninda/notifier/pkg"

	"github.com/spf13/cobra"
)

var sendEmail = &cobra.Command{
	Use:   "sendEmail",
	Short: "Sends email",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.SendEmail(cmd)
	},
}

func init() {
	sendEmail.PersistentFlags().StringP("to", "t", "", "Destination email")
	sendEmail.PersistentFlags().StringP("subject", "", "", "Email subject")
	sendEmail.PersistentFlags().StringP("body", "", "", "Email body")
	sendEmail.PersistentFlags().BoolP("disable-tls", "", false, "Disable StartTLS")

}
