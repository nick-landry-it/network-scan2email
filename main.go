package main

import (
	"fmt"
	"ipupdater/cryptography"
	"ipupdater/mailer"
	"ipupdater/networking"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func main() {
	// WIP
	encryptedEmail, err := os.ReadFile("./encryptedfiles/.encryptedEmail.lck")
	check(err)
	encryptedPassword, err := os.ReadFile("./encryptedfiles/.encryptedPassword.lck")
	check(err)
	encryptionKey, err := os.ReadFile("./encryptedfiles/.encryptionKey.lck")
	check(err)

	ipPlaceholder := networking.GetNetwork()
	fmt.Println(ipPlaceholder)
	decryptedEmail := cryptography.Decrypt(string(encryptedEmail), string(encryptionKey))
	decryptedPassword := cryptography.Decrypt(string(encryptedPassword), string(encryptionKey))
	fmt.Printf("Decrypted Email : %s\n", decryptedEmail)
	fmt.Printf("Decrypted Password : %s\n", decryptedPassword)

	mailer.SendMail(decryptedEmail, decryptedPassword)
}
