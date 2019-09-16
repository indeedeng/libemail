package libemail

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// MessageMock implements Message
type MessageMock struct {
	t minimock.Tester

	funcCompile          func() (ba1 []byte, err error)
	inspectFuncCompile   func()
	afterCompileCounter  uint64
	beforeCompileCounter uint64
	CompileMock          mMessageMockCompile

	funcRecipients          func() (sa1 []string)
	inspectFuncRecipients   func()
	afterRecipientsCounter  uint64
	beforeRecipientsCounter uint64
	RecipientsMock          mMessageMockRecipients

	funcSender          func() (s1 string)
	inspectFuncSender   func()
	afterSenderCounter  uint64
	beforeSenderCounter uint64
	SenderMock          mMessageMockSender
}

// NewMessageMock returns a mock for Message
func NewMessageMock(t minimock.Tester) *MessageMock {
	m := &MessageMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CompileMock = mMessageMockCompile{mock: m}

	m.RecipientsMock = mMessageMockRecipients{mock: m}

	m.SenderMock = mMessageMockSender{mock: m}

	return m
}

type mMessageMockCompile struct {
	mock               *MessageMock
	defaultExpectation *MessageMockCompileExpectation
	expectations       []*MessageMockCompileExpectation
}

// MessageMockCompileExpectation specifies expectation struct of the Message.Compile
type MessageMockCompileExpectation struct {
	mock *MessageMock

	results *MessageMockCompileResults
	Counter uint64
}

// MessageMockCompileResults contains results of the Message.Compile
type MessageMockCompileResults struct {
	ba1 []byte
	err error
}

// Expect sets up expected params for Message.Compile
func (mmCompile *mMessageMockCompile) Expect() *mMessageMockCompile {
	if mmCompile.mock.funcCompile != nil {
		mmCompile.mock.t.Fatalf("MessageMock.Compile mock is already set by Set")
	}

	if mmCompile.defaultExpectation == nil {
		mmCompile.defaultExpectation = &MessageMockCompileExpectation{}
	}

	return mmCompile
}

// Inspect accepts an inspector function that has same arguments as the Message.Compile
func (mmCompile *mMessageMockCompile) Inspect(f func()) *mMessageMockCompile {
	if mmCompile.mock.inspectFuncCompile != nil {
		mmCompile.mock.t.Fatalf("Inspect function is already set for MessageMock.Compile")
	}

	mmCompile.mock.inspectFuncCompile = f

	return mmCompile
}

// Return sets up results that will be returned by Message.Compile
func (mmCompile *mMessageMockCompile) Return(ba1 []byte, err error) *MessageMock {
	if mmCompile.mock.funcCompile != nil {
		mmCompile.mock.t.Fatalf("MessageMock.Compile mock is already set by Set")
	}

	if mmCompile.defaultExpectation == nil {
		mmCompile.defaultExpectation = &MessageMockCompileExpectation{mock: mmCompile.mock}
	}
	mmCompile.defaultExpectation.results = &MessageMockCompileResults{ba1, err}
	return mmCompile.mock
}

//Set uses given function f to mock the Message.Compile method
func (mmCompile *mMessageMockCompile) Set(f func() (ba1 []byte, err error)) *MessageMock {
	if mmCompile.defaultExpectation != nil {
		mmCompile.mock.t.Fatalf("Default expectation is already set for the Message.Compile method")
	}

	if len(mmCompile.expectations) > 0 {
		mmCompile.mock.t.Fatalf("Some expectations are already set for the Message.Compile method")
	}

	mmCompile.mock.funcCompile = f
	return mmCompile.mock
}

// Compile implements Message
func (mmCompile *MessageMock) Compile() (ba1 []byte, err error) {
	mm_atomic.AddUint64(&mmCompile.beforeCompileCounter, 1)
	defer mm_atomic.AddUint64(&mmCompile.afterCompileCounter, 1)

	if mmCompile.inspectFuncCompile != nil {
		mmCompile.inspectFuncCompile()
	}

	if mmCompile.CompileMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCompile.CompileMock.defaultExpectation.Counter, 1)

		mm_results := mmCompile.CompileMock.defaultExpectation.results
		if mm_results == nil {
			mmCompile.t.Fatal("No results are set for the MessageMock.Compile")
		}
		return (*mm_results).ba1, (*mm_results).err
	}
	if mmCompile.funcCompile != nil {
		return mmCompile.funcCompile()
	}
	mmCompile.t.Fatalf("Unexpected call to MessageMock.Compile.")
	return
}

// CompileAfterCounter returns a count of finished MessageMock.Compile invocations
func (mmCompile *MessageMock) CompileAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCompile.afterCompileCounter)
}

// CompileBeforeCounter returns a count of MessageMock.Compile invocations
func (mmCompile *MessageMock) CompileBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCompile.beforeCompileCounter)
}

// MinimockCompileDone returns true if the count of the Compile invocations corresponds
// the number of defined expectations
func (m *MessageMock) MinimockCompileDone() bool {
	for _, e := range m.CompileMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CompileMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCompileCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCompile != nil && mm_atomic.LoadUint64(&m.afterCompileCounter) < 1 {
		return false
	}
	return true
}

// MinimockCompileInspect logs each unmet expectation
func (m *MessageMock) MinimockCompileInspect() {
	for _, e := range m.CompileMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MessageMock.Compile")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CompileMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCompileCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Compile")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCompile != nil && mm_atomic.LoadUint64(&m.afterCompileCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Compile")
	}
}

type mMessageMockRecipients struct {
	mock               *MessageMock
	defaultExpectation *MessageMockRecipientsExpectation
	expectations       []*MessageMockRecipientsExpectation
}

// MessageMockRecipientsExpectation specifies expectation struct of the Message.Recipients
type MessageMockRecipientsExpectation struct {
	mock *MessageMock

	results *MessageMockRecipientsResults
	Counter uint64
}

// MessageMockRecipientsResults contains results of the Message.Recipients
type MessageMockRecipientsResults struct {
	sa1 []string
}

// Expect sets up expected params for Message.Recipients
func (mmRecipients *mMessageMockRecipients) Expect() *mMessageMockRecipients {
	if mmRecipients.mock.funcRecipients != nil {
		mmRecipients.mock.t.Fatalf("MessageMock.Recipients mock is already set by Set")
	}

	if mmRecipients.defaultExpectation == nil {
		mmRecipients.defaultExpectation = &MessageMockRecipientsExpectation{}
	}

	return mmRecipients
}

// Inspect accepts an inspector function that has same arguments as the Message.Recipients
func (mmRecipients *mMessageMockRecipients) Inspect(f func()) *mMessageMockRecipients {
	if mmRecipients.mock.inspectFuncRecipients != nil {
		mmRecipients.mock.t.Fatalf("Inspect function is already set for MessageMock.Recipients")
	}

	mmRecipients.mock.inspectFuncRecipients = f

	return mmRecipients
}

// Return sets up results that will be returned by Message.Recipients
func (mmRecipients *mMessageMockRecipients) Return(sa1 []string) *MessageMock {
	if mmRecipients.mock.funcRecipients != nil {
		mmRecipients.mock.t.Fatalf("MessageMock.Recipients mock is already set by Set")
	}

	if mmRecipients.defaultExpectation == nil {
		mmRecipients.defaultExpectation = &MessageMockRecipientsExpectation{mock: mmRecipients.mock}
	}
	mmRecipients.defaultExpectation.results = &MessageMockRecipientsResults{sa1}
	return mmRecipients.mock
}

//Set uses given function f to mock the Message.Recipients method
func (mmRecipients *mMessageMockRecipients) Set(f func() (sa1 []string)) *MessageMock {
	if mmRecipients.defaultExpectation != nil {
		mmRecipients.mock.t.Fatalf("Default expectation is already set for the Message.Recipients method")
	}

	if len(mmRecipients.expectations) > 0 {
		mmRecipients.mock.t.Fatalf("Some expectations are already set for the Message.Recipients method")
	}

	mmRecipients.mock.funcRecipients = f
	return mmRecipients.mock
}

// Recipients implements Message
func (mmRecipients *MessageMock) Recipients() (sa1 []string) {
	mm_atomic.AddUint64(&mmRecipients.beforeRecipientsCounter, 1)
	defer mm_atomic.AddUint64(&mmRecipients.afterRecipientsCounter, 1)

	if mmRecipients.inspectFuncRecipients != nil {
		mmRecipients.inspectFuncRecipients()
	}

	if mmRecipients.RecipientsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRecipients.RecipientsMock.defaultExpectation.Counter, 1)

		mm_results := mmRecipients.RecipientsMock.defaultExpectation.results
		if mm_results == nil {
			mmRecipients.t.Fatal("No results are set for the MessageMock.Recipients")
		}
		return (*mm_results).sa1
	}
	if mmRecipients.funcRecipients != nil {
		return mmRecipients.funcRecipients()
	}
	mmRecipients.t.Fatalf("Unexpected call to MessageMock.Recipients.")
	return
}

// RecipientsAfterCounter returns a count of finished MessageMock.Recipients invocations
func (mmRecipients *MessageMock) RecipientsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRecipients.afterRecipientsCounter)
}

// RecipientsBeforeCounter returns a count of MessageMock.Recipients invocations
func (mmRecipients *MessageMock) RecipientsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRecipients.beforeRecipientsCounter)
}

// MinimockRecipientsDone returns true if the count of the Recipients invocations corresponds
// the number of defined expectations
func (m *MessageMock) MinimockRecipientsDone() bool {
	for _, e := range m.RecipientsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RecipientsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRecipientsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRecipients != nil && mm_atomic.LoadUint64(&m.afterRecipientsCounter) < 1 {
		return false
	}
	return true
}

// MinimockRecipientsInspect logs each unmet expectation
func (m *MessageMock) MinimockRecipientsInspect() {
	for _, e := range m.RecipientsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MessageMock.Recipients")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RecipientsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRecipientsCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Recipients")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRecipients != nil && mm_atomic.LoadUint64(&m.afterRecipientsCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Recipients")
	}
}

type mMessageMockSender struct {
	mock               *MessageMock
	defaultExpectation *MessageMockSenderExpectation
	expectations       []*MessageMockSenderExpectation
}

// MessageMockSenderExpectation specifies expectation struct of the Message.Sender
type MessageMockSenderExpectation struct {
	mock *MessageMock

	results *MessageMockSenderResults
	Counter uint64
}

// MessageMockSenderResults contains results of the Message.Sender
type MessageMockSenderResults struct {
	s1 string
}

// Expect sets up expected params for Message.Sender
func (mmSender *mMessageMockSender) Expect() *mMessageMockSender {
	if mmSender.mock.funcSender != nil {
		mmSender.mock.t.Fatalf("MessageMock.Sender mock is already set by Set")
	}

	if mmSender.defaultExpectation == nil {
		mmSender.defaultExpectation = &MessageMockSenderExpectation{}
	}

	return mmSender
}

// Inspect accepts an inspector function that has same arguments as the Message.Sender
func (mmSender *mMessageMockSender) Inspect(f func()) *mMessageMockSender {
	if mmSender.mock.inspectFuncSender != nil {
		mmSender.mock.t.Fatalf("Inspect function is already set for MessageMock.Sender")
	}

	mmSender.mock.inspectFuncSender = f

	return mmSender
}

// Return sets up results that will be returned by Message.Sender
func (mmSender *mMessageMockSender) Return(s1 string) *MessageMock {
	if mmSender.mock.funcSender != nil {
		mmSender.mock.t.Fatalf("MessageMock.Sender mock is already set by Set")
	}

	if mmSender.defaultExpectation == nil {
		mmSender.defaultExpectation = &MessageMockSenderExpectation{mock: mmSender.mock}
	}
	mmSender.defaultExpectation.results = &MessageMockSenderResults{s1}
	return mmSender.mock
}

//Set uses given function f to mock the Message.Sender method
func (mmSender *mMessageMockSender) Set(f func() (s1 string)) *MessageMock {
	if mmSender.defaultExpectation != nil {
		mmSender.mock.t.Fatalf("Default expectation is already set for the Message.Sender method")
	}

	if len(mmSender.expectations) > 0 {
		mmSender.mock.t.Fatalf("Some expectations are already set for the Message.Sender method")
	}

	mmSender.mock.funcSender = f
	return mmSender.mock
}

// Sender implements Message
func (mmSender *MessageMock) Sender() (s1 string) {
	mm_atomic.AddUint64(&mmSender.beforeSenderCounter, 1)
	defer mm_atomic.AddUint64(&mmSender.afterSenderCounter, 1)

	if mmSender.inspectFuncSender != nil {
		mmSender.inspectFuncSender()
	}

	if mmSender.SenderMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSender.SenderMock.defaultExpectation.Counter, 1)

		mm_results := mmSender.SenderMock.defaultExpectation.results
		if mm_results == nil {
			mmSender.t.Fatal("No results are set for the MessageMock.Sender")
		}
		return (*mm_results).s1
	}
	if mmSender.funcSender != nil {
		return mmSender.funcSender()
	}
	mmSender.t.Fatalf("Unexpected call to MessageMock.Sender.")
	return
}

// SenderAfterCounter returns a count of finished MessageMock.Sender invocations
func (mmSender *MessageMock) SenderAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSender.afterSenderCounter)
}

// SenderBeforeCounter returns a count of MessageMock.Sender invocations
func (mmSender *MessageMock) SenderBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSender.beforeSenderCounter)
}

// MinimockSenderDone returns true if the count of the Sender invocations corresponds
// the number of defined expectations
func (m *MessageMock) MinimockSenderDone() bool {
	for _, e := range m.SenderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SenderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSenderCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSender != nil && mm_atomic.LoadUint64(&m.afterSenderCounter) < 1 {
		return false
	}
	return true
}

// MinimockSenderInspect logs each unmet expectation
func (m *MessageMock) MinimockSenderInspect() {
	for _, e := range m.SenderMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MessageMock.Sender")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SenderMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSenderCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Sender")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSender != nil && mm_atomic.LoadUint64(&m.afterSenderCounter) < 1 {
		m.t.Error("Expected call to MessageMock.Sender")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCompileInspect()

		m.MinimockRecipientsInspect()

		m.MinimockSenderInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MessageMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCompileDone() &&
		m.MinimockRecipientsDone() &&
		m.MinimockSenderDone()
}
