package emailutil

import (
	"fmt"
	"strings"

	"github.com/go-gomail/gomail"
)

type MailUser struct {
	Email string
	Name  string
}

const (
	MailServerHostExmail = "smtp.exmail.qq.com"
	MailServerHost126    = "smtp.126.com"
	MailServerHost163    = "smtp.163.com"
)

var (
	mailServerHost = MailServerHostExmail
	mailNickname   = "no-reply"
	mailUsername   = "no-reply@airdb.net"
	mailPassword   = ""
	mailServerPort = 465
)

var maintainers = []MailUser{
	{
		Email: "airdb@qq.com",
		Name:  "airdb",
	},
}

func SendEmail(toEmails, subject, content string) {
	m := gomail.NewMessage()

	ccList := []string{}
	for _, maintainer := range maintainers {
		ccList = append(ccList, m.FormatAddress(maintainer.Email, maintainer.Name))
	}

	m.SetAddressHeader("From", mailUsername, mailNickname)

	toList := []string{}
	for _, toEmail := range strings.Split(toEmails, ",") {
		toList = append(toList, m.FormatAddress(toEmail, ""))
	}
	// m.SetHeader("To", m.FormatAddress(toMail.Email, toMail.Name))
	m.SetHeader("To", toList...)

	m.SetHeader("Cc", ccList...)

	m.SetHeader("Subject", subject)

	m.SetBody("text/html", content)

	d := gomail.NewDialer(mailServerHost, mailServerPort, mailUsername, mailPassword)
	fmt.Println("dd", d)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("err", err.Error())
		return
	}
}
