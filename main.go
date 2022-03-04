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
	// WIP Reading in Encrypted Files
	encryptedEmail, err := os.ReadFile("./encryptedfiles/.encryptedEmail.lck")
	check(err)
	encryptedPassword, err := os.ReadFile("./encryptedfiles/.encryptedPassword.lck")
	check(err)
	encryptionKey, err := os.ReadFile("./encryptedfiles/.encryptionKey.lck")
	check(err)

	// Grab the External IP using the GetNetwork function
	// and print to console to make sure it works
	// brainstorming ideas on sending IP only when it changes
	currentIP := networking.GetNetwork()
	sendableIP := currentIP
	fmt.Println(currentIP + " -> " + sendableIP)

	// Utilize the Cryptography library Decrypt function
	// and print to console to make sure it works
	decryptedEmail := cryptography.Decrypt(string(encryptedEmail), string(encryptionKey))
	decryptedPassword := cryptography.Decrypt(string(encryptedPassword), string(encryptionKey))
	fmt.Printf("Decrypted Email : %s\n", decryptedEmail)
	fmt.Printf("Decrypted Password : %s\n", decryptedPassword)

	// SendMail function requires (email, password, body)
	mailer.SendMail(decryptedEmail, decryptedPassword, sendableIP)
}
