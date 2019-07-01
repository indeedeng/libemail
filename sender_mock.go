package libemail

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// SenderMock implements Sender
type SenderMock struct {
	t minimock.Tester

	funcSend          func(m1 Message) (err error)
	afterSendCounter  uint64
	beforeSendCounter uint64
	SendMock          mSenderMockSend
}

// NewSenderMock returns a mock for Sender
func NewSenderMock(t minimock.Tester) *SenderMock {
	m := &SenderMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendMock = mSenderMockSend{mock: m}
	m.SendMock.callArgs = []*SenderMockSendParams{}

	return m
}

type mSenderMockSend struct {
	mock               *SenderMock
	defaultExpectation *SenderMockSendExpectation
	expectations       []*SenderMockSendExpectation

	callArgs []*SenderMockSendParams
	mutex    sync.RWMutex
}

// SenderMockSendExpectation specifies expectation struct of the Sender.Send
type SenderMockSendExpectation struct {
	mock    *SenderMock
	params  *SenderMockSendParams
	results *SenderMockSendResults
	Counter uint64
}

// SenderMockSendParams contains parameters of the Sender.Send
type SenderMockSendParams struct {
	m1 Message
}

// SenderMockSendResults contains results of the Sender.Send
type SenderMockSendResults struct {
	err error
}

// Expect sets up expected params for Sender.Send
func (mmSend *mSenderMockSend) Expect(m1 Message) *mSenderMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("SenderMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &SenderMockSendExpectation{}
	}

	mmSend.defaultExpectation.params = &SenderMockSendParams{m1}
	for _, e := range mmSend.expectations {
		if minimock.Equal(e.params, mmSend.defaultExpectation.params) {
			mmSend.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSend.defaultExpectation.params)
		}
	}

	return mmSend
}

// Return sets up results that will be returned by Sender.Send
func (mmSend *mSenderMockSend) Return(err error) *SenderMock {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("SenderMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &SenderMockSendExpectation{mock: mmSend.mock}
	}
	mmSend.defaultExpectation.results = &SenderMockSendResults{err}
	return mmSend.mock
}

//Set uses given function f to mock the Sender.Send method
func (mmSend *mSenderMockSend) Set(f func(m1 Message) (err error)) *SenderMock {
	if mmSend.defaultExpectation != nil {
		mmSend.mock.t.Fatalf("Default expectation is already set for the Sender.Send method")
	}

	if len(mmSend.expectations) > 0 {
		mmSend.mock.t.Fatalf("Some expectations are already set for the Sender.Send method")
	}

	mmSend.mock.funcSend = f
	return mmSend.mock
}

// When sets expectation for the Sender.Send which will trigger the result defined by the following
// Then helper
func (mmSend *mSenderMockSend) When(m1 Message) *SenderMockSendExpectation {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("SenderMock.Send mock is already set by Set")
	}

	expectation := &SenderMockSendExpectation{
		mock:   mmSend.mock,
		params: &SenderMockSendParams{m1},
	}
	mmSend.expectations = append(mmSend.expectations, expectation)
	return expectation
}

// Then sets up Sender.Send return parameters for the expectation previously defined by the When method
func (e *SenderMockSendExpectation) Then(err error) *SenderMock {
	e.results = &SenderMockSendResults{err}
	return e.mock
}

// Send implements Sender
func (mmSend *SenderMock) Send(m1 Message) (err error) {
	mm_atomic.AddUint64(&mmSend.beforeSendCounter, 1)
	defer mm_atomic.AddUint64(&mmSend.afterSendCounter, 1)

	params := &SenderMockSendParams{m1}

	// Record call args
	mmSend.SendMock.mutex.Lock()
	mmSend.SendMock.callArgs = append(mmSend.SendMock.callArgs, params)
	mmSend.SendMock.mutex.Unlock()

	for _, e := range mmSend.SendMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSend.SendMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSend.SendMock.defaultExpectation.Counter, 1)
		want := mmSend.SendMock.defaultExpectation.params
		got := SenderMockSendParams{m1}
		if want != nil && !minimock.Equal(*want, got) {
			mmSend.t.Errorf("SenderMock.Send got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSend.SendMock.defaultExpectation.results
		if results == nil {
			mmSend.t.Fatal("No results are set for the SenderMock.Send")
		}
		return (*results).err
	}
	if mmSend.funcSend != nil {
		return mmSend.funcSend(m1)
	}
	mmSend.t.Fatalf("Unexpected call to SenderMock.Send. %v", m1)
	return
}

// SendAfterCounter returns a count of finished SenderMock.Send invocations
func (mmSend *SenderMock) SendAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.afterSendCounter)
}

// SendBeforeCounter returns a count of SenderMock.Send invocations
func (mmSend *SenderMock) SendBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.beforeSendCounter)
}

// Calls returns a list of arguments used in each call to SenderMock.Send.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSend *mSenderMockSend) Calls() []*SenderMockSendParams {
	mmSend.mutex.RLock()

	argCopy := make([]*SenderMockSendParams, len(mmSend.callArgs))
	copy(argCopy, mmSend.callArgs)

	mmSend.mutex.RUnlock()

	return argCopy
}

// MinimockSendDone returns true if the count of the Send invocations corresponds
// the number of defined expectations
func (m *SenderMock) MinimockSendDone() bool {
	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSend != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendInspect logs each unmet expectation
func (m *SenderMock) MinimockSendInspect() {
	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SenderMock.Send with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		if m.SendMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SenderMock.Send")
		} else {
			m.t.Errorf("Expected call to SenderMock.Send with params: %#v", *m.SendMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSend != nil && mm_atomic.LoadUint64(&m.afterSendCounter) < 1 {
		m.t.Error("Expected call to SenderMock.Send")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SenderMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSendInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SenderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *SenderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendDone()
}
