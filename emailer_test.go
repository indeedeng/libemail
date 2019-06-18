package libemail

import (
	"errors"
	"net/smtp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSMTPSender_no_address(t *testing.T) {
	sendMailFunc := func(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
		return nil
	}

	require.Panics(t, func() {
		_ = NewSMTPSender(SMTPSenderOptions{
			Address:      "",
			SendMailFunc: sendMailFunc,
			Auth:         nil,
		})
	})
}

func TestNewSMTPSender_no_func(t *testing.T) {
	require.Panics(t, func() {
		_ = NewSMTPSender(SMTPSenderOptions{
			Address:      "127.0.0.1:12345",
			SendMailFunc: nil,
			Auth:         nil,
		})
	})
}

func Test_NewSMTPSender_Send(t *testing.T) {
	executed := false
	sendMailFunc := func(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
		executed = true
		return nil
	}

	s := NewSMTPSender(SMTPSenderOptions{
		Address:      "127.0.0.1:12345",
		SendMailFunc: sendMailFunc,
		Auth:         nil,
	})

	err := s.Send(&TextEmail{
		From:        "alice@example.com",
		To:          []string{"bob@example.com"},
		ReplyTo:     "rt@example.com",
		Subject:     "my subject",
		Body:        "hello world",
		Attachments: nil,
	})

	require.NoError(t, err)
	require.True(t, executed)
}

func Test_NewSMTPSender_Send_no_compile(t *testing.T) {
	messageMock := NewMessageMock(t)
	defer messageMock.MinimockFinish()

	messageMock.CompileMock.Return(nil, errors.New("bad message"))

	executed := false
	sendMailFunc := func(addr string, auth smtp.Auth, from string, to []string, msg []byte) error {
		executed = true
		return nil
	}

	s := NewSMTPSender(SMTPSenderOptions{
		Address:      "127.0.0.1:12345",
		SendMailFunc: sendMailFunc,
		Auth:         nil,
	})

	err := s.Send(messageMock)

	require.Error(t, err)
	require.False(t, executed)
}
