// Package libemail provides utilities for composing plaintext or HTML email
// messages with attachments that can be sent through a typical SMTP server.
package libemail // import "oss.indeed.com/go/libemail"

import (
	"net/smtp"
)

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -g -i Sender -s _mock.go

// Sender is an interface to be implemented by something which is able to
// Send a Message.
type Sender interface {
	Send(Message) error
}

// SMTPSender is an implementation of Sender that defers to the standard
// library smtp package for sending a Message.
type SMTPSender struct {
	address      string
	sendMailFunc SendMailFunc
	auth         SMTPAuth
}

// SMTPSenderOptions are used to configure an SMTPSender.
type SMTPSenderOptions struct {
	Address      string       // e.g. localhost:10025 for postfix
	SendMailFunc SendMailFunc // e.g. smtp.SendMail
	Auth         SMTPAuth     // e.g. smtp.Auth, or nil
}

// NewSMTPSender creates an SMTPSender configured with
// the given options.
//
// The Address and SendMailFunc fields are required, and will
// cause a panic if left as zero values. The Auth fields can be
// left nil, if the underlying SMTP server does not have auth enabled.
func NewSMTPSender(options SMTPSenderOptions) Sender {
	if options.Address == "" {
		panic("libemail: Address must be configured")
	}

	if options.SendMailFunc == nil {
		panic("libemail: SendMailFunc must be configured")
	}

	return &SMTPSender{
		address:      options.Address,
		sendMailFunc: options.SendMailFunc,
		auth:         options.Auth,
	}
}

func (s *SMTPSender) Send(m Message) error {
	msgBytes, err := m.Compile()
	if err != nil {
		return err
	}

	return s.sendMailFunc(
		s.address,
		s.auth,
		m.Sender(),
		m.Recipients(),
		msgBytes,
	)
}

// SendMailFunc would typically be implemented by simply calling smtp.SendMail
// but is left as a configurable function for testing purposes.
type SendMailFunc func(addr string, a smtp.Auth, from string, to []string, msg []byte) error

//go:generate go run github.com/gojuno/minimock/v3/cmd/minimock -g -i SMTPAuth -s _mock.go

// SMTPAuth enables generating mock for smtp.Auth
type SMTPAuth interface {
	smtp.Auth
}
