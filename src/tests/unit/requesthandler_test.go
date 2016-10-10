package unit

import (
	"errors"
	"testing"

	"github.com/masud328np/go_ansible_docker_aws/src/lib"
	"github.com/stretchr/testify/mock"
)

func Test_HandleCallRead(t *testing.T) {
	handler := lib.NewHandler()
	mockConn := new(MockConnection)
	msgBytes := []byte("HI")
	mockConn.On("Read", mock.Anything).Run(func(args mock.Arguments) {
		args[0] = msgBytes
	}).Return(len(msgBytes), errors.New(""))
	mockConn.On("Close").Return(nil)
	handler.Handle(mockConn)

	mockConn.AssertCalled(t, "Read", mock.Anything)
}

func Test_HandleIfMsgRespond(t *testing.T) {
	handler := lib.NewHandler()
	mockConn := new(MockConnection)
	buf := make([]byte, 1024)
	msgBytes := []byte("HI")
	mockConn.On("Read", buf).Return(12, nil)

	mockConn.On("Close").Return(nil)
	mockConn.On("Write", mock.Anything).Return(len(msgBytes), nil)

	handler.Handle(mockConn)

	mockConn.AssertCalled(t, "Write", mock.Anything)
}
