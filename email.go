package email

import (
    "crypto/tls"
    "gopkg.in/gomail.v2"
)



func SendEmail(sender string, recipients []string, subject string, message string, message_type string, attachments []string) {
    m := gomail.NewMessage()
    m.SetHeader("From", sender)
    m.SetHeader("To", recipients...)
    m.SetHeader("Subject", subject)
    if message_type == "html" {
        m.SetBody("text/html", message)
    } else {
        m.SetBody("text/plain", message)
    }
    if len(attachments) > 0 {
	for _, file := range attachments {
            m.Attach(file)
        }
    }

    d := gomail.Dialer{Host: "localhost", Port: 25}
    d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    // Send the email to Bob, Cora and Dan.
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

