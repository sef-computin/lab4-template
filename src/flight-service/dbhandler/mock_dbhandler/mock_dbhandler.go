// Code generated by MockGen. DO NOT EDIT.
// Source: dbhandler.go

// Package mock_dbhandler is a generated GoMock package.
package mock_dbhandler

import (
        models "lab2/src/flight-service/models"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockFlightDB is a mock of FlightDB interface.
type MockFlightDB struct {
        ctrl     *gomock.Controller
        recorder *MockFlightDBMockRecorder
}

// MockFlightDBMockRecorder is the mock recorder for MockFlightDB.
type MockFlightDBMockRecorder struct {
        mock *MockFlightDB
}

// NewMockFlightDB creates a new mock instance.
func NewMockFlightDB(ctrl *gomock.Controller) *MockFlightDB {
        mock := &MockFlightDB{ctrl: ctrl}
        mock.recorder = &MockFlightDBMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlightDB) EXPECT() *MockFlightDBMockRecorder {
        return m.recorder
}

// GetAirportByID mocks base method.
func (m *MockFlightDB) GetAirportByID(airportID string) (*models.Airport, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAirportByID", airportID)
        ret0, _ := ret[0].(*models.Airport)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetAirportByID indicates an expected call of GetAirportByID.
func (mr *MockFlightDBMockRecorder) GetAirportByID(airportID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAirportByID", reflect.TypeOf((*MockFlightDB)(nil).GetAirportByID), airportID)
}

// GetAllFlights mocks base method.
func (m *MockFlightDB) GetAllFlights() ([]*models.Flight, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAllFlights")
        ret0, _ := ret[0].([]*models.Flight)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetAllFlights indicates an expected call of GetAllFlights.
func (mr *MockFlightDBMockRecorder) GetAllFlights() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFlights", reflect.TypeOf((*MockFlightDB)(nil).GetAllFlights))
}

// GetFlightByNumber mocks base method.
func (m *MockFlightDB) GetFlightByNumber(flightNumber string) (*models.Flight, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetFlightByNumber", flightNumber)
        ret0, _ := ret[0].(*models.Flight)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetFlightByNumber indicates an expected call of GetFlightByNumber.
func (mr *MockFlightDBMockRecorder) GetFlightByNumber(flightNumber interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlightByNumber", reflect.TypeOf((*MockFlightDB)(nil).GetFlightByNumber), flightNumber)
}