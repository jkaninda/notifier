package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"io/ioutil"
	"jkaninda/notifier/util"
	"net/http"
	"os"
)

func getUrl(token string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", token)

}

func SendMessage(cmd *cobra.Command) {
	//Load env
	err := godotenv.Load()
	if err != nil {
		util.Info(".env file not found")
		//util.Error("Error loading .env file, file not found")
	}
	message = util.GetEnv(cmd, "message", "TG_MESSAGE")
	token := os.Getenv("TG_TOKEN")
	chatId := os.Getenv("TG_CHAT_ID")

	var vars = []string{
		"TG_TOKEN",
		"TG_CHAT_ID",
	}
	err = util.CheckEnvVars(vars)
	if err != nil {
		util.Fatal("Required environment variables needed, %v", err)
	}
	body, _ := json.Marshal(map[string]string{
		"chat_id": chatId,
		"text":    message,
	})
	url := fmt.Sprintf("%s/sendMessage", getUrl(token))

	// Create a HTTP post request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	code := response.StatusCode
	if code == 200 {
		util.Info("Message has been sent")
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		util.Fatal("Message not sent, error: %s", string(body))
	}

}
