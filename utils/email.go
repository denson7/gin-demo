package utils

import (
	"fmt"
	"net/smtp"
)

func SendEmail()  {
	// Sender data.
	from := "2671475697@qq.com"
	password := "liudongsheng727"

	// Receiver email address.
	to := []string{
		"denson.liu@greenpacket.com.cn",
	}

	// smtp server configuration.
	smtpHost := "smtp.qq.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}