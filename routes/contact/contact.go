package contact

import (
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
)

type postContactDetails struct {
	Username string `json:"username,omitempty"`
	Message  string `json:"message,omitempty"`
}

func PostContactHandler(w http.ResponseWriter, r *http.Request) {
	var data postContactDetails
	json.NewDecoder(r.Body).Decode(&data)

	from := "test@deonsummers.com"
	pass := "Test1234"
	to := "deonsummers01@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		data.Message

	err := smtp.SendMail("smtp.mail.us-west-2.awsapps.com:465",
		smtp.PlainAuth("", from, pass, "smtp.mail.us-west-2.awsapps.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}
