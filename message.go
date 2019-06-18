package libemail

import "errors"

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i Message -s _mock.go

// A Message represents something which can be formatted into an RFC-compatible
// email document, including headers, text, html, and attachments.
type Message interface {
	Sender() string
	Recipients() []string
	Compile() ([]byte, error)
}

var (
	ErrAttachmentAlreadySet = errors.New("attachment already exists")
	ErrMissingFromField     = errors.New("'from' field not set")
	ErrMissingToField       = errors.New("'to' field not set")
)

const (
	boundary                = "z4kq10mn34xlw0302dlwkld492dk"
	contentType             = "Content-Type"
	contentTransferEncoding = "Content-Transfer-Encoding"
	contentDisposition      = "Content-Disposition"
	applicationOctetStream  = "application/octet-stream"
	base64Encoding          = "base64"
)

func requiredFieldsSet(m Message) error {
	if m.Sender() == "" {
		return ErrMissingFromField
	}

	if len(m.Recipients()) == 0 {
		return ErrMissingToField
	}

	return nil
}
