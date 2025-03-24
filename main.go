package main

import (
	"log"
	"net/smtp"

	lua "github.com/yuin/gopher-lua"
)

// sendEmail sends an email using Gmail's SMTP server
func sendEmail(L *lua.LState) int {
	to := L.ToString(1)
	subject := L.ToString(2)
	body := L.ToString(3)

	// set these env variables
	from := "xxxxxxxxx"
	password := "xxxxxxx"

	if from == "" || password == "" {
		L.Push(lua.LString("Error: SMTP_EMAIL or SMTP_PASSWORD is not set"))
		return 1
	}

	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth("", from, password, smtpServer)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(smtpServer+":587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		L.Push(lua.LString("Error: " + err.Error()))
		return 1
	}

	L.Push(lua.LString("Email sent successfully"))
	return 1
}

func main() {
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("send_email", L.NewFunction(sendEmail))

	err := L.DoFile("temp_check.lua")
	if err != nil {
		log.Fatal(err)
	}
}
