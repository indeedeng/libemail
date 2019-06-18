package libemail

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HTMLEmail_String(t *testing.T) {
	email := &HTMLEmail{
		From:     "noreply@example.com",
		To:       []string{"devnull@example.com"},
		Subject:  "test email",
		HtmlBody: "<strong>test email</strong>",
		TextBody: "test email text",
	}

	// "from: noreply@example.com, to: [devnull@example.com], body: <strong>test email</strong>"
	s := email.String()
	require.True(t, strings.Contains(s, "from: noreply@example.com"))
	require.True(t, strings.Contains(s, "to: [devnull@example.com]"))
	require.True(t, strings.Contains(s, "htmlbody: <strong>test email</strong>"))
	require.True(t, strings.Contains(s, "textbody: test email text"))
}

func Test_HTMLEmail_Sender(t *testing.T) {
	email := &HTMLEmail{
		From: "noreply@example.com",
	}

	sender := email.Sender()
	require.Equal(t, "noreply@example.com", sender)
}

func Test_HTMLEmail_Recipients(t *testing.T) {
	email := &HTMLEmail{
		To: []string{
			"a@example.com",
			"b@example.com",
		},
	}

	recipients := email.Recipients()
	require.Contains(t, recipients, "a@example.com")
	require.Contains(t, recipients, "b@example.com")
}

func Test_HTMLEmail_Compile(t *testing.T) {
	email := &HTMLEmail{
		From:     "noreply@example.com",
		To:       []string{"devnull@example.com"},
		Subject:  "test email",
		HtmlBody: "<strong>test email</strong>",
		TextBody: "test email text",
	}

	bs, err := email.Compile()
	require.NoError(t, err)
	require.True(t, len(bs) > 0)
}

func Test_HTMLEmail_Compile_noTo(t *testing.T) {
	email := &HTMLEmail{
		From:     "noreply@example.com",
		To:       []string{},
		Subject:  "test email",
		HtmlBody: "<strong>test email</strong>",
		TextBody: "test email text",
	}

	_, err := email.Compile()
	require.EqualError(t, err, ErrMissingToField.Error())
}

func Test_HTMLEmail_Compile_noFrom(t *testing.T) {
	email := &HTMLEmail{
		From:     "",
		To:       []string{"a@example.com", "b@example.com"},
		Subject:  "test email",
		HtmlBody: "<strong>test email</strong>",
		TextBody: "test email text",
	}

	_, err := email.Compile()
	require.EqualError(t, err, ErrMissingFromField.Error())
}
