package cmd

import (
	"github.com/jkaninda/notifier/pkg"

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
	sendMail.PersistentFlags().StringP("to", "t", "", "Comma-separated list of email addresses")
	sendMail.PersistentFlags().StringP("subject", "", "", "Email subject")
	sendMail.PersistentFlags().StringP("body", "", "", "Email body")
	sendMail.PersistentFlags().BoolP("insecure", "", false, "Disable StartTLS")
	sendMail.PersistentFlags().StringP("attach", "", "", "Attach a file")
}
