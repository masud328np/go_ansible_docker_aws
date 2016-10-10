package unit

import (
	"errors"
	"testing"

	"github.com/masud328np/go_ansible_docker_aws/src/lib"
	"github.com/stretchr/testify/mock"
)

func Test_StartlisteningIfCallListenSuccessThenTrue(t *testing.T) {

	mockNet := new(MockNetWrapper)
	mockListener := new(MockListener)
	mockHandler := new(MockHandler)
	mockNet.On("Listen", mock.Anything, mock.Anything).Return(mockListener, nil)
	mockListener.On("Close").Return(nil)
	mockListener.On("Accept").Return(new(MockConnection), errors.New(""))
	mockHandler.On("Handle", mock.Anything)
	echoserver := lib.NewServer(mockHandler, mockNet)

	result := echoserver.StartListening("8090")

	mockNet.AssertCalled(t, "Listen", mock.Anything, mock.Anything)
	mockListener.AssertCalled(t, "Close")
	if result == false {
		t.Fail()
	}
}

func Test_StartlisteningIfCallListenReturnFalseThenfalse(t *testing.T) {

	mockNet := new(MockNetWrapper)
	mockListener := new(MockListener)

	mockNet.On("Listen", mock.Anything, mock.Anything).Return(mockListener, errors.New(""))
	echoserver := lib.NewServer(nil, mockNet)

	result := echoserver.StartListening("8090")

	mockNet.AssertCalled(t, "Listen", mock.Anything, mock.Anything)
	if result != false {
		t.Fail()
	}
}

func Test_IfListenerSuccessCallAccept(t *testing.T) {

	mockNet := new(MockNetWrapper)
	mockListener := new(MockListener)
	mockConn := new(MockConnection)
	mockHandler := new(MockHandler)
	mockNet.On("Listen", mock.Anything, mock.Anything).Return(mockListener, nil)
	mockListener.On("Accept").Return(mockConn, errors.New(""))
	mockListener.On("Close").Return(nil)
	mockHandler.On("Handle", mock.Anything)
	echoserver := lib.NewServer(mockHandler, mockNet)

	result := echoserver.StartListening("8090")

	mockNet.AssertCalled(t, "Listen", mock.Anything, mock.Anything)
	mockListener.AssertCalled(t, "Accept")
	if result == false {
		t.Fail()
	}
}

func Test_IfAcceptConnectionCallEchoHandler(t *testing.T) {

	mockNet := new(MockNetWrapper)
	mockListener := new(MockListener)
	mockConn := new(MockConnection)
	mockHandler := new(MockHandler)

	mockNet.On("Listen", mock.Anything, mock.Anything).Return(mockListener, nil)
	mockListener.On("Accept").Return(mockConn, errors.New(""))
	mockListener.On("Close").Return(nil)
	mockHandler.On("Handle", mockConn)

	echoserver := lib.NewServer(mockHandler, mockNet)

	echoserver.StartListening("8090")

	mockNet.AssertCalled(t, "Listen", mock.Anything, mock.Anything)
	mockListener.AssertCalled(t, "Accept")
	mockHandler.AssertCalled(t, "Handle", mockConn)

}
