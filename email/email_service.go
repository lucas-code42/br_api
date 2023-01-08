package email

import "net/smtp"

func SendEmail() {
	from := "lucasdearaujo.brandao@gmail.com"
	password := "dtxwgzdlvngoctfg"

	toEmailAddress := from
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Te amo meu amor\n"
	body := "Enviandio email pelo golang cheeeeck"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}

}
