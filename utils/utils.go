package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_defi/utils/constants"
	"log"
	"net/http"
	"strings"
)

func PrintDashed() {
	fmt.Println(strings.Repeat("-", 75))
}

func SendTelegramMessage(text string) {
	values := map[string]string{"chat_id": constants.TelegramChatId, "text": text, "pase_mode": "HTML"}

	jsonValue, _ := json.Marshal(values)
	req, err := http.Post(constants.TelegramUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()
}
