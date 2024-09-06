package pkg

import (
	"crypto/tls"
	"strings"

	"github.com/jkaninda/notifier/util"

	"github.com/go-mail/mail"
	"github.com/spf13/cobra"
)

func SendEmail(cmd *cobra.Command) {
	util.Info("Start sending email....")
	config := getMailConfig(cmd)
	emails := strings.Split(config.mailTo, ",")
	m := mail.NewMessage()
	m.SetHeader("From", config.mailFrom)
	m.SetHeader("To", emails...)
	if config.attachFile != "" {
		m.Attach(config.attachFile)
	}
	m.SetHeader("Subject", config.mailSubject)
	m.SetBody("text/html", config.mailBody)
	d := mail.NewDialer(config.mailHost, config.mailPort, config.mailUserName, config.mailPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: config.skipTls}

	if err := d.DialAndSend(m); err != nil {
		util.Fatal("Error could not send email : %v", err)
	}
	util.Info("Email has been sent")

}
