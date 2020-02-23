package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type bptMsg struct {
	Contact      string `json:"contact"`
	Text         string `json:"text"`
	Number       string `json:"number"`
	Name_bot     string `json:"name_bot"`
	Number_check string `json:"number_check"`
	Site_name    string `json:"site_name"`
}
type ansWer struct {
	Result bool   `json:"result"`
	Text   string `json:"text"`
}

func sendBot(w http.ResponseWriter, r *http.Request) {

	botList := map[string]string{"siuzanna": "@siuzanna_j"}

	w.Header().Set("Content-Type", "application/json")

	bm := &bptMsg{}
	err := json.NewDecoder(r.Body).Decode(bm) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		aw := ansWer{Result: false,
			Text: "Что то пошло нетак. Не можем отправить сообщение. Свяжитесь со мной по email."}
		json.NewEncoder(w).Encode(&aw)
		return
	}

	if len(bm.Number) < 1 {
		aw := ansWer{Result: false,
			Text: "Поле с примером пустое."}
		json.NewEncoder(w).Encode(&aw)

		return
	}
	bm_Number, err := strconv.Atoi(bm.Number)
	if err != nil {
		aw := ansWer{Result: false,
			Text: "Что то пошло нетак. Не можем отправить сообщение. Свяжитесь со мной по email."}
		json.NewEncoder(w).Encode(&aw)
		return
	}
	if bm_Number < 10 {
		aw := ansWer{Result: false,
			Text: "Ты не правильно решил пример."}
		json.NewEncoder(w).Encode(&aw)

		return
	}
	if bm.Number != bm.Number_check {
		aw := ansWer{Result: false,
			Text: "Ты чего надо же ввести число " + bm.Number_check}
		json.NewEncoder(w).Encode(&aw)

		return
	}
	if len(bm.Contact) < 4 {
		aw := ansWer{Result: false,
			Text: "Поле с Email/Telegram/Телефон должно быть по длине больше 4 символов"}
		json.NewEncoder(w).Encode(&aw)

		return
	}
	if len(bm.Text) < 4 {
		aw := ansWer{Result: false,
			Text: "Поле Краткая суть должен быть конечно кратким но больше 4 символов уж точно"}
		json.NewEncoder(w).Encode(&aw)

		return
	}
	if vaf, ok := botList[bm.Name_bot]; ok {
		urltelegram := "https://api.telegram.org/bot1044261349:AAHHaVhxRqoghSIcHVxWmR1UHrTfxfAZEAk/sendMessage?chat_id=" +
			vaf + "&text=" + url.QueryEscape("Сайт:\n    "+bm.Site_name+"\nКонтакт:\n    "+bm.Contact+"\nСообщение:    "+bm.Text)
		_, err := http.Get(urltelegram)
		if err != nil {
			aw := ansWer{Result: false,
				Text: "Что то  пошло не так с отправкой сообщения. Пожалуйста позвоните мне"}
			json.NewEncoder(w).Encode(&aw)
			return
		}
		aw := ansWer{Result: true,
			Text: "Все супер. Свяжусь с вами каr только освобожусь."}
		json.NewEncoder(w).Encode(&aw)
		return
	} else {
		aw := ansWer{Result: false,
			Text: "Что то пошло нетак. Не можем отправить сообщение. Свяжитесь со мной по email."}
		json.NewEncoder(w).Encode(&aw)
		return
	}
	aw := ansWer{Result: false,
		Text: "!!!!"}
	json.NewEncoder(w).Encode(&aw)
	return

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/botik", sendBot).Methods("POST")
	log.Fatal(http.ListenAndServe(":8005", r))
}
