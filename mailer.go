package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/wneessen/go-mail"
)

type Credentials struct {
	Email string
	Pass  string
	Smtp  string
	Port  int
}

const credentialsPath = "./credentials.json"

var credentials *Credentials

func init() {
	creds, err := readCredentials(credentialsPath)
	credentials = creds
	if err != nil {
		fmt.Println("There was an error reading the credentials, the app won't send any email.", err)
	}

}

func SendMail(message string, to string, subject string) error {

	if credentials == nil {
		return nil
	}

	creds := credentials

	m := mail.NewMsg()
	err := m.From(creds.Email)
	if err != nil {
		return fmt.Errorf("error setting 'from': %w", err)
	}
	err = m.To(to)
	if err != nil {
		return fmt.Errorf("error setting 'to': %w", err)
	}
	m.Subject(subject)

	m.SetBodyString("text/html", message)

	d, err := mail.NewClient(creds.Smtp, mail.WithPort(creds.Port), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithPassword(creds.Pass), mail.WithUsername(creds.Email))
	if err != nil {
		return err
	}

	err = d.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}
	defer fmt.Println("Succesfully sent email to destinatary: ", to)
	return nil
}

func readCredentials(filePath string) (*Credentials, error) {

	if credentials != nil {
		return credentials, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var obj Credentials
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil

}

func FormatErrorEmail(url string, message string) string {
	styletag := `<style>
		span {
			font-family:roboto;
			font-size: 1rem;
			color: green;
		}
		a{
			font-size: inherit;
			color: blue;
		}
		h1 {
			font-size: 1.5rem;
			color: black;
		}
		p{
			color:red;
			font-family:monospace;
			font-size: 1rem;
		}
	</style>`

	date := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf(`%s
	<span>%s</span>
	<h1>The url <a href"%s">%s</a> is down with the following error:</h1>
	<p>%s</p>
	`, styletag, date, url, url, strings.ReplaceAll(message, "\n", "<br>"))

}

func FormatUpEmail(url string) string {
	styletag := `<style>
	span {
		font-family:roboto;
		font-size: 1rem;
		color: green;
	}
	a{
		font-size: inherit;
		color: blue;
	}
	h1 {
		font-size: 1.5rem;
		color: black;
	}
</style>`

	date := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf(`%s
<span>%s</span>
<h1>The url <a href"%s">%s</a> is is up again</h1>
`, styletag, date, url, url)

}
