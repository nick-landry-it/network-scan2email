package mailer

func SendMail(from string, password string) (string, string) {
	//toList := []string{"example@gmail.com"}
	//host := "smtp.gmail.com"
	//port := "587"
	//msg := "Hello, World!"
	//body := []byte(msg)
	//auth := smtp.PlainAuth("", from, password, host)
	//err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//fmt.Println("Successfully sent mail to all users in toList")

	return from, password
}
