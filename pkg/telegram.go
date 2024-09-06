package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jkaninda/notifier/util"

	"github.com/spf13/cobra"
)

func SendMessage(cmd *cobra.Command) {

	config := getTgConfig(cmd)

	body, _ := json.Marshal(map[string]string{
		"chat_id": config.chatId,
		"text":    config.message,
	})
	url := fmt.Sprintf("%s/sendMessage", getTgUrl(config.token))

	// Create an HTTP post request
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
