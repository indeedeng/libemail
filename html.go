package libemail

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

// HTMLEmail is a representation of a html encoded email. You can
// specify the sender, a list of recipients, the subject, and an
// html encoded body as well as a default utf-8 plain text body.
type HTMLEmail struct {
	From     string
	To       []string
	ReplyTo  string
	Subject  string
	HtmlBody string
	TextBody string
}

// String returns a human readable representation of a HTMLEmail.
func (e *HTMLEmail) String() string {
	s := fmt.Sprintf(
		"from: %s, to: %s, htmlbody: %s; textbody: %s",
		e.From,
		e.To,
		e.HtmlBody,
		e.TextBody,
	)

	return s
}

func (e *HTMLEmail) Sender() string {
	return e.From
}

func (e *HTMLEmail) Recipients() []string {
	// make a copy of the To slice, so the original is
	// not mutated if the caller wants to modify the slice they get
	dst := make([]string, len(e.To))
	copy(dst, e.To)
	return dst
}

func (e *HTMLEmail) Compile() ([]byte, error) {
	if err := requiredFieldsSet(e); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	buf.WriteString("From: " + e.From + "\n")
	now := time.Now()
	buf.WriteString("Date: " + now.Format(time.RFC822) + "\n")
	buf.WriteString("To: " + strings.Join(e.To, ",") + "\n")
	if e.ReplyTo != "" {
		buf.WriteString("Reply-To: " + e.ReplyTo + "\n")
	}
	buf.WriteString("Subject: " + e.Subject + "\n")
	buf.WriteString("Content-Type: multipart/alternative; boundary=\"" + boundary + "\"\n")
	buf.WriteString("MIME-Version: 1.0\n\n")
	buf.WriteString("--" + boundary)
	buf.WriteString("\n")
	if e.TextBody != "" {
		// if the text body isn't empty, lets put it in here
		buf.WriteString("Content-Type: text/plain; charset=utf-8\n\n")
		buf.WriteString(e.TextBody)
		buf.WriteString("\n")
		buf.WriteString("--" + boundary)
		buf.WriteString("\n")
	}

	// utf8 text body of the email
	buf.WriteString("Content-type: text/html; charset=utf-8\n\n")
	buf.WriteString(e.HtmlBody)
	buf.WriteString("\n")
	buf.WriteString("--" + boundary + "--")

	return buf.Bytes(), nil
}
