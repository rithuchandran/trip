package mock_service

import (
	"github.com/rithuchandran/trip/domain"
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockTripServiceInt is a mock of TripServiceInt interface
type MockTripServiceInt struct {
	ctrl     *gomock.Controller
	recorder *MockTripServiceIntMockRecorder
}

// MockTripServiceIntMockRecorder is the mock recorder for MockTripServiceInt
type MockTripServiceIntMockRecorder struct {
	mock *MockTripServiceInt
}

// NewMockTripServiceInt creates a new mock instance
func NewMockTripServiceInt(ctrl *gomock.Controller) *MockTripServiceInt {
	mock := &MockTripServiceInt{ctrl: ctrl}
	mock.recorder = &MockTripServiceIntMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTripServiceInt) EXPECT() *MockTripServiceIntMockRecorder {
	return m.recorder
}

// CreateTrip mocks base method
func (m *MockTripServiceInt) CreateTrip(arg0 domain.Trip) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateTrip", arg0)
}

// CreateTrip indicates an expected call of CreateTrip
func (mr *MockTripServiceIntMockRecorder) CreateTrip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrip", reflect.TypeOf((*MockTripServiceInt)(nil).CreateTrip), arg0)
}

// DeleteTrip mocks base method
func (m *MockTripServiceInt) DeleteTrip(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrip", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrip indicates an expected call of DeleteTrip
func (mr *MockTripServiceIntMockRecorder) DeleteTrip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrip", reflect.TypeOf((*MockTripServiceInt)(nil).DeleteTrip), arg0)
}

// GetTrip mocks base method
func (m *MockTripServiceInt) GetTrip(arg0 int) (domain.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrip", arg0)
	ret0, _ := ret[0].(domain.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrip indicates an expected call of GetTrip
func (mr *MockTripServiceIntMockRecorder) GetTrip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrip", reflect.TypeOf((*MockTripServiceInt)(nil).GetTrip), arg0)
}

// UpdateTrip mocks base method
func (m *MockTripServiceInt) UpdateTrip(arg0 int, arg1 domain.Trip) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateTrip", arg0, arg1)
}

// UpdateTrip indicates an expected call of UpdateTrip
func (mr *MockTripServiceIntMockRecorder) UpdateTrip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrip", reflect.TypeOf((*MockTripServiceInt)(nil).UpdateTrip), arg0, arg1)
}
