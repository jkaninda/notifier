package pkg

import (
	"crypto/tls"
	"jkaninda/notifier/util"
	"os"

	"github.com/go-mail/mail"
	"github.com/spf13/cobra"
)

func SendEmail(cmd *cobra.Command) {
	//Load env
	envFile, _ = cmd.Flags().GetString("env-file")
	util.LoadEnv(envFile)

	mailHost = os.Getenv("MAIL_HOST")
	mailPort = util.GetIntEnv("MAIL_PORT")
	mailUserName = os.Getenv("MAIL_USERNAME")
	mailPassword = os.Getenv("MAIL_PASSWORD")
	mailTo = util.GetEnv(cmd, "to", "MAIL_TO")
	mailSubject = util.GetEnv(cmd, "subject", "MAIL_SUBJECT")
	mailBody = util.GetEnv(cmd, "message", "MAIL_BODY")
	mailFrom = os.Getenv("MAIL_FROM")
	skipTls, _ = cmd.Flags().GetBool("insecure")
	attachFile, _ = cmd.Flags().GetString("attach")

	util.Info("Start sending email....")
	var vars = []string{
		"MAIL_HOST",
		"MAIL_USERNAME",
		"MAIL_PASSWORD",
		"MAIL_TO",
		"MAIL_SUBJECT",
		"MAIL_FROM",
	}
	err := util.CheckEnvVars(vars)
	if err != nil {
		util.Fatal("Required environment variables needed, %v", err)
	}
	m := mail.NewMessage()
	m.SetHeader("From", mailFrom)
	m.SetHeader("To", mailTo)
	if attachFile != "" {
		m.Attach(attachFile)
	}
	//m.SetAddressHeader("Cc", "jonas@mailhog.local", "Jonas")
	m.SetHeader("Subject", mailSubject)
	m.SetBody("text/html", mailBody)

	d := mail.NewDialer(mailHost, mailPort, mailUserName, mailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		util.Fatal("Error could not send email : %v", err)
	}
	util.Info("Email has been sent")

}
