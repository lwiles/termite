// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package run

import (
	"sync"
)

// Ensure, that ExecutorMock does implement Executor.
// If this is not the case, regenerate this file with moq.
var _ Executor = &ExecutorMock{}

// ExecutorMock is a mock implementation of Executor.
//
//	func TestSomethingThatUsesExecutor(t *testing.T) {
//
//		// make and configure a mocked Executor
//		mockedExecutor := &ExecutorMock{
//			OutputFunc: func(cmd *Cmd) ([]byte, error) {
//				panic("mock out the Output method")
//			},
//			RunFunc: func(cmd *Cmd) error {
//				panic("mock out the Run method")
//			},
//		}
//
//		// use mockedExecutor in code that requires Executor
//		// and then make assertions.
//
//	}
type ExecutorMock struct {
	// OutputFunc mocks the Output method.
	OutputFunc func(cmd *Cmd) ([]byte, error)

	// RunFunc mocks the Run method.
	RunFunc func(cmd *Cmd) error

	// calls tracks calls to the methods.
	calls struct {
		// Output holds details about calls to the Output method.
		Output []struct {
			// Cmd is the cmd argument value.
			Cmd *Cmd
		}
		// Run holds details about calls to the Run method.
		Run []struct {
			// Cmd is the cmd argument value.
			Cmd *Cmd
		}
	}
	lockOutput sync.RWMutex
	lockRun    sync.RWMutex
}

// Output calls OutputFunc.
func (mock *ExecutorMock) Output(cmd *Cmd) ([]byte, error) {
	if mock.OutputFunc == nil {
		panic("ExecutorMock.OutputFunc: method is nil but Executor.Output was just called")
	}
	callInfo := struct {
		Cmd *Cmd
	}{
		Cmd: cmd,
	}
	mock.lockOutput.Lock()
	mock.calls.Output = append(mock.calls.Output, callInfo)
	mock.lockOutput.Unlock()
	return mock.OutputFunc(cmd)
}

// OutputCalls gets all the calls that were made to Output.
// Check the length with:
//
//	len(mockedExecutor.OutputCalls())
func (mock *ExecutorMock) OutputCalls() []struct {
	Cmd *Cmd
} {
	var calls []struct {
		Cmd *Cmd
	}
	mock.lockOutput.RLock()
	calls = mock.calls.Output
	mock.lockOutput.RUnlock()
	return calls
}

// Run calls RunFunc.
func (mock *ExecutorMock) Run(cmd *Cmd) error {
	if mock.RunFunc == nil {
		panic("ExecutorMock.RunFunc: method is nil but Executor.Run was just called")
	}
	callInfo := struct {
		Cmd *Cmd
	}{
		Cmd: cmd,
	}
	mock.lockRun.Lock()
	mock.calls.Run = append(mock.calls.Run, callInfo)
	mock.lockRun.Unlock()
	return mock.RunFunc(cmd)
}

// RunCalls gets all the calls that were made to Run.
// Check the length with:
//
//	len(mockedExecutor.RunCalls())
func (mock *ExecutorMock) RunCalls() []struct {
	Cmd *Cmd
} {
	var calls []struct {
		Cmd *Cmd
	}
	mock.lockRun.RLock()
	calls = mock.calls.Run
	mock.lockRun.RUnlock()
	return calls
}