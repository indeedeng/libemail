package libemail

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TextEmail_String(t *testing.T) {
	a1 := strings.NewReader("some kind of attachment")
	a2 := strings.NewReader("some other attachment")
	email := &TextEmail{
		From:    "noreply@example.com",
		To:      []string{"devnull@example.com"},
		Subject: "test email",
		Body:    "this is the body of a test email",
	}
	require.NoError(t, email.Attach("a1.txt", a1))
	require.NoError(t, email.Attach("a2.txt", a2))

	// "from: noreply@example.com, to: [devnull@example.com], body: this is the body of a test email, attachments: [a2.txt a1.txt]"
	s := email.String()
	require.True(t, strings.Contains(s, "from: noreply@example.com"))
	require.True(t, strings.Contains(s, "to: [devnull@example.com]"))
	require.True(t, strings.Contains(s, "body: this is the body of a test email"))
	require.True(t, strings.Contains(s, "a1.txt"))
	require.True(t, strings.Contains(s, "a2.txt"))
}

func Test_TextEmail_Sender(t *testing.T) {
	email := &TextEmail{
		From: "noreply@example.com",
	}
	sender := email.Sender()
	require.Equal(t, "noreply@example.com", sender)
}

func Test_TextEmail_Recipients(t *testing.T) {
	email := &TextEmail{
		To: []string{"a@example.com", "b@example.com"},
	}
	recipients := email.Recipients()
	require.Contains(t, recipients, "a@example.com")
	require.Contains(t, recipients, "b@example.com")
}

func Test_TextEmail_Attachment_duplicate(t *testing.T) {
	a1 := strings.NewReader("some kind of attachment")
	a2 := strings.NewReader("some other attachment")
	email := &TextEmail{
		From:    "noreply@example.com",
		To:      []string{"devnull@example.com"},
		Subject: "test email",
		Body:    "this is the body of a test email",
	}
	require.NoError(t, email.Attach("a1.txt", a1))

	err := email.Attach("a1.txt", a2)
	require.EqualError(t, err, ErrAttachmentAlreadySet.Error())
}

func Test_TextEmail_Compile(t *testing.T) {
	a1 := strings.NewReader("some kind of attachment")
	a2 := strings.NewReader("some other attachment")

	email := &TextEmail{
		From:    "noreply@example.com",
		To:      []string{"devnull@example.com"},
		Subject: "test email",
		Body:    "this is the body of a test email",
	}

	require.NoError(t, email.Attach("a1.txt", a1))
	require.NoError(t, email.Attach("a2.txt", a2))

	_, err := email.Compile()
	require.NoError(t, err)
}
