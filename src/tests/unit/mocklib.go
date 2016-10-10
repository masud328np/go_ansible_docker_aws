package unit

import (
	"net"

	"github.com/stretchr/testify/mock"
)

//mocking area

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(conn net.Conn) {
	m.Called(conn)
	return
}

type MockConnection struct {
	net.Conn
	mock.Mock
}

func (m *MockConnection) Read(input []byte) (int, error) {
	args := m.Called(input)

	if args[1] == nil {
		return args[0].(int), nil
	}
	return args[0].(int), args[1].(error)
}

func (m *MockConnection) Write(input []byte) (int, error) {
	args := m.Called(input)

	if args[1] == nil {
		return args[0].(int), nil
	}
	return args[0].(int), args[1].(error)
}

func (m *MockConnection) Close() error {
	return nil
}

type MockListener struct {
	net.Listener
	mock.Mock
}

func (m *MockListener) Accept() (net.Conn, error) {
	args := m.Called()
	if args[1] == nil {
		return args[0].(net.Conn), nil
	}
	return args[0].(net.Conn), args[1].(error)
}

func (m *MockListener) Close() error {
	args := m.Called()
	if args[0] == nil {
		return nil
	}
	return args[0].(error)
}

type MockNetWrapper struct {
	mock.Mock
}

func (m *MockNetWrapper) Listen(proto string, addrPort string) (net.Listener, error) {
	args := m.Called(proto, addrPort)
	if args[1] == nil {
		return args[0].(net.Listener), nil
	}
	return args[0].(net.Listener), args[1].(error)
}
