package alert

import (
	"fmt"
	"io"
	"net/http"
)

func SendMsg(msg string) {

	url := "https://api.telegram.org/bot6865573955:AAGTfiloEPbsqZrKsM23YDD1kKghLH0wHa0/sendMessage?chat_id=-4115383073&text=" + msg
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("send msg ：" + err.Error())
		return
	}
}
