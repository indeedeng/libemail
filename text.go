package libemail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/textproto"
	"strings"
	"time"
)

// TextEmail is a representation of a plaintext email. You can
// specify the sender, a list of recipients, the subject,
// the body, a set of attachments.
type TextEmail struct {
	From        string
	To          []string
	ReplyTo     string
	Subject     string
	Body        string
	Attachments map[string]*Attachment
}

// Attachment is a representation of an email attachment.
// There is currently no support for inline attachments.
type Attachment struct {
	Filename string
	Content  []byte
	Header   textproto.MIMEHeader
}

// String returns a human readable representation of a TextEmail.
func (e *TextEmail) String() string {
	s := fmt.Sprintf(
		"from: %s, to: %s, body: %s",
		e.From,
		e.To,
		e.Body,
	)

	if len(e.Attachments) > 0 {
		attachments := make([]string, 0, len(e.Attachments))
		for _, attachment := range e.Attachments {
			attachments = append(attachments, attachment.Filename)
		}
		s = fmt.Sprintf("%s, attachments: %s", s, attachments)
	}

	return s
}

// Sender returns the configured address of the sender of the TextEmail.
func (e *TextEmail) Sender() string {
	return e.From
}

// Recipients returns the complete list of recipients. Since
// TextEmail currently does not support CC or BCC, this is the same as
// the "To" field.
func (e *TextEmail) Recipients() []string {
	// make a copy of the To slice, so the original is
	// not mutated if the caller wants to modify the slice they get
	dst := make([]string, len(e.To))
	copy(dst, e.To)
	return dst
}

// Attach adds creates an Attachment with filename and an io.Reader full
// of the attachment content.
//
// Attach can be called multiple times to append more than one document.
//
// If two attachments with the same filename are added, an error is returned.
//
// An error is returned if the content of an attachment could not be read.
func (e *TextEmail) Attach(filename string, content io.Reader) error {
	if e.Attachments == nil {
		e.Attachments = make(map[string]*Attachment)
	}

	if _, exists := e.Attachments[filename]; exists {
		return ErrAttachmentAlreadySet
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, content); err != nil {
		return err
	}
	a := &Attachment{
		Filename: filename,
		Header:   textproto.MIMEHeader{},
		Content:  buf.Bytes(),
	}
	a.Header.Set(contentType, applicationOctetStream)
	a.Header.Set(contentTransferEncoding, base64Encoding)
	a.Header.Set(contentDisposition, fmt.Sprintf("attachment;\r\n filename=\"%s\"", filename))
	e.Attachments[filename] = a
	return nil
}

// Compile returns turns the TextEmail into an RFC compatible set of bytes.
// If some required field is not present, an error is returned.
func (e *TextEmail) Compile() ([]byte, error) {
	if err := requiredFieldsSet(e); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	buf.WriteString("From: " + e.From + "\r\n")
	now := time.Now()
	buf.WriteString("Date: " + now.Format(time.RFC822) + "\r\n")
	buf.WriteString("To: " + strings.Join(e.To, ", \r\n ") + "\r\n")
	if e.ReplyTo != "" {
		buf.WriteString("Reply-To: " + e.ReplyTo + "\r\n")
	}
	buf.WriteString("Subject: " + e.Subject + "\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")

	// multipart with boundary if e has attachments
	if len(e.Attachments) > 0 {
		buf.WriteString(contentType + ": multipart/mixed; boundary=" + boundary + "\r\n")
		buf.WriteString("--" + boundary + "\r\n")
	}

	// utf8 text body of the email
	buf.WriteString(contentType + ": text/plain; charset=utf-8\r\n\r\n")
	buf.WriteString(e.Body)
	buf.WriteString("\r\n")

	// add attachments if any
	for _, a := range e.Attachments {
		buf.WriteString("\r\n\r\n--" + boundary + "\r\n")
		buf.WriteString(contentType + ": " + a.Header.Get(contentType) + "\r\n")
		buf.WriteString(contentTransferEncoding + ": " + a.Header.Get(contentTransferEncoding) + "\r\n")
		buf.WriteString(contentDisposition + ": " + a.Header.Get(contentDisposition) + "\r\n")
		encoded := make([]byte, base64.StdEncoding.EncodedLen(len(a.Content)))
		base64.StdEncoding.Encode(encoded, a.Content)
		buf.Write(encoded)
		buf.WriteString("\r\n--" + boundary)
	}
	if len(e.Attachments) > 0 {
		buf.WriteString("--")
	}

	return buf.Bytes(), nil
}
