libemail
========

[![Go Report Card](https://goreportcard.com/badge/oss.indeed.com/go/libemail)](https://goreportcard.com/report/oss.indeed.com/go/libemail)
[![Build Status](https://travis-ci.com/indeedeng/libemail.svg?branch=master)](https://travis-ci.com/indeedeng/libemail)
[![GoDoc](https://godoc.org/oss.indeed.com/go/libemail?status.svg)](https://godoc.org/oss.indeed.com/go/libemail)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/indeedeng/libemail.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/indeedeng/libemail.svg)](LICENSE)

# Project Overview

Package `libemail` provides utilities for composing plaintext or HTML email
messages with attachments that can be sent through a typical SMTP server.

# Getting Started

The `libemail` module can be installed like any other Go module, e.g.
```
$ go get oss.indeed.com/go/libemail
```

This is a typical example for sending a basic plaintext email to a locally
running SMTP server, e.g.
```
m := &libemail.TextEmail{
    From:    "alice@example.com",
    To:      []string{"bob@example.com"},
    ReplyTo: "alice@example.com",
    Subject: "this is a test",
    Body:    "hello world!",
}

smtpSender := libemail.NewSMTPSender(libemail.SMTPSenderOptions{
    Address:      "localhost:10025",
    SendMailFunc: smtp.SendMail,
    Auth:         nil,
})

_ = smtpSender.Send(m)
```

# Asking Questions

For technical questions about `libemail`, just file an issue in the GitHub tracker.

For questions about Open Source in Indeed Engineering, send us an email at
opensource@indeed.com

# Contributing

We welcome contributions! Feel free to help make `libemail` better.

### Process

- Open an issue and describe the desired feature / bug fix before making
changes. It's useful to get a second pair of eyes before investing development
effort.
- Make the change. If adding a new feature, remember to provide tests that
demonstrate the new feature works, including any error paths. If contributing
a bug fix, add tests that demonstrate the erroneous behavior is fixed.
- Open a pull request. Automated CI tests will run. If the tests fail, please
make changes to fix the behavior, and repeat until the tests pass.
- Once everything looks good, one of the indeedeng members will review the
PR and provide feedback.

# Maintainers

The `oss.indeed.com/go/libemail` module is maintained by Indeed Engineering.

While we are always busy helping people get jobs, we will try to respond to
GitHub issues, pull requests, and questions within a couple of business days.

# Code of Conduct

`oss.indeed.com/go/libemail` is governed by the [Contributer Covenant v1.4.1](CODE_OF_CONDUCT.md)

For more information please contact opensource@indeed.com.

# License

The `oss.indeed.com/go/libemail` module is open source under the [BSD-3-Clause](LICENSE)
license.
