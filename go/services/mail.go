package services

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendUserEmail(email string, token string, newUser bool) bool {
	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	app_host := os.Getenv("APP_HOST")

	emailString := strings.Replace(email, "@", "%40", -1)

	message := fmt.Sprintf("ユーザー登録を進める場合は<a href=\"%s/user/create/?email=%s&token=%s\">こちら</a>をクリックしてください。", app_host, emailString, token)

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_ADMIN"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "メールアドレスの認証")
	m.SetBody("text/html", message)
	m.AddAlternative("text/plain", "メールアドレスの認証")

	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return true
}
