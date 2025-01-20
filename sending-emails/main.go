package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the API key from the environment variables
	gmail_apiKey := os.Getenv("GMAIL_API_KEY")
	if gmail_apiKey == "" {
		log.Fatal("API_KEY is not set in the .env file")
	}

	// Create a new message
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", "indexpage.me@gmail.com")
	message.SetHeader("To", "isaquefrankli@gmail.com")
	message.SetHeader("Subject", "Hello from the Golang team")

	// Set email body
	message.SetBody("text/plain", "This is the Test Body")

	// Set up the SMTP dialer
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "indexpage.me@gmail.com", gmail_apiKey)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
