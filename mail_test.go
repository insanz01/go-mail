package mailservice_test

import (
	"os"
	"testing"

	mail "mail-service"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	err := godotenv.Load()
	require.NoError(t, err)

	// Baca variabel lingkungan dari file .env
	senderName := os.Getenv("EMAIL_SENDER_NAME")
	senderAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	senderPassword := os.Getenv("EMAIL_SENDER_PASSWORD")

	sender := mail.NewGmailSender(senderName, senderAddress, senderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="insandev.com">Insan Dev</a></p>`

	to := []string{"insankamil002@gmail.com", "insankamil.uad@gmail.com"}
	attachFiles := []string{"ketupat.png"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
