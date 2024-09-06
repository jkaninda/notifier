package pkg

import (
	"fmt"
	"os"
	"strings"

	"github.com/jkaninda/notifier/util"
	"github.com/spf13/cobra"
)

type tgConfig struct {
	message string
	token   string
	chatId  string
}
type mailConfig struct {
	mailHost     string
	mailPort     int
	mailUserName string
	mailPassword string
	mailTo       string
	mailSubject  string
	mailBody     string
	mailFrom     string
	skipTls      bool
	attachFile   string
}

func getMailConfig(cmd *cobra.Command) *mailConfig {
	//Load env
	envFile, _ = cmd.Flags().GetString("env-file")
	util.LoadEnv(envFile)
	var vars = []string{
		"MAIL_HOST",
		"MAIL_USERNAME",
		"MAIL_PASSWORD",
		"MAIL_TO",
		"MAIL_SUBJECT",
		"MAIL_FROM",
	}
	mail := mailConfig{}
	mail.mailHost = os.Getenv("MAIL_HOST")
	mail.mailPort = util.GetIntEnv("MAIL_PORT")
	mail.mailUserName = os.Getenv("MAIL_USERNAME")
	mail.mailPassword = os.Getenv("MAIL_PASSWORD")

	mail.mailSubject = util.GetEnv(cmd, "subject", "MAIL_SUBJECT")
	mail.mailBody = util.GetEnv(cmd, "body", "MAIL_BODY")
	mail.mailFrom = os.Getenv("MAIL_FROM")
	mail.skipTls, _ = cmd.Flags().GetBool("insecure")
	mail.attachFile, _ = cmd.Flags().GetString("attach")

	emails := strings.Split(util.GetEnv(cmd, "to", "MAIL_TO"), ",")
	for i, email := range emails {
		emails[i] = fmt.Sprintf("\"%s\"", email)
	}
	mail.mailTo = util.GetEnv(cmd, "to", "MAIL_TO")
	err := util.CheckEnvVars(vars)
	if err != nil {
		util.Fatal("Required environment variables for Mail needed, %v", err)
	}
	return &mail

}
func getTgConfig(cmd *cobra.Command) *tgConfig {
	envFile, _ = cmd.Flags().GetString("env-file")
	util.LoadEnv(envFile)

	var vars = []string{
		"TG_TOKEN",
		"TG_CHAT_ID",
	}

	tg := tgConfig{}
	tg.message = util.GetEnv(cmd, "message", "TG_MESSAGE")
	tg.token = os.Getenv("TG_TOKEN")
	tg.chatId = os.Getenv("TG_CHAT_ID")
	err := util.CheckEnvVars(vars)
	if err != nil {
		util.Fatal("Required environment variables for Telegram needed, %v", err)
	}
	return &tg
}

func getTgUrl(token string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", token)

}
