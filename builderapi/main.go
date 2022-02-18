package main

import (
	"fmt"
	"os"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		fmt.Errorf("invalid email syntax, correct example: name@domain.com")
		os.Exit(1)
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendEmailImpl(email *email) {
	// actual sends the email
	fmt.Println(email)
}

type build func(*EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmailImpl(&builder.email)
}

func main() {
	// SendEmail call
	SendEmail(func(b *EmailBuilder) {
		b.From("me@domain.com").To("you@domain.com").
			Subject("Meeting").Body("Hello You, time to go for a coffe break")
	})
}
