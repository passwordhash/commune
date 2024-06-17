package service

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	D EmailDeps
}

func NewEmailService(deps EmailDeps) *EmailService {
	return &EmailService{
		deps,
	}
}

func (s *EmailService) SendCode(to string, passcode string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.D.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Commune | код-пароль")
	m.SetBody("text/html", s.getEmailTemplate(passcode))

	d := gomail.NewDialer(s.D.SmtpHost, s.D.SmtpPort, s.D.From, s.D.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d.DialAndSend(m)
}

func (s *EmailService) getEmailTemplate(passcode string) string {
	return `
	<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Код подтверждения - Commune</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f8f9fa;
            color: #343a40;
            margin: 0;
            padding: 0;
        }
        .container {
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
            text-align: center;
            background-color: #ffffff;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        .icon {
            font-size: 48px;
            color: rgb(76, 175, 80);
            margin-bottom: 20px;
        }
        h1 {
            margin-bottom: 20px;
            font-size: 24px;
            color: #4caf50;
        }
        p {
            font-size: 16px;
            color: #6c757d;
        }
        .code {
            font-size: 20px;
            font-weight: bold;
            margin: 20px 0;
            padding: 10px;
            border: 1px dashed #4caf50;
            display: inline-block;
            color: #4caf50;
        }
        .footer {
            margin-top: 20px;
            padding-top: 10px;
            border-top: 1px solid #dee2e6;
            color: #6c757d;
        }
        .footer a {
            color: #6c757d;
            text-decoration: none;
        }
        .footer a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <i class="fas fa-comments icon"></i>
        <h1>Commune - общайтесь без границ</h1>
        <p>Ваш код для входа в аккаунт:</p>
        <div class="code">` + passcode + `</div>
        <p>Введите этот код на сайте, чтобы продолжить.</p>
    </div>
    <div class="footer">
        <p>Сделано автором</p>
        <a href="https://github.com/passwordhash" target="_blank"><i class="fab fa-github"></i> Github</a>
    </div>
</body>
</html>

`
}
