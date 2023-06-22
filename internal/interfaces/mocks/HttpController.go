// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HttpController is an autogenerated mock type for the HttpController type
type HttpController struct {
	mock.Mock
}

// CancelRecording provides a mock function with given fields: writer, request
func (_m *HttpController) CancelRecording(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// CancelReplay provides a mock function with given fields: writer, request
func (_m *HttpController) CancelReplay(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// ExportRecordedData provides a mock function with given fields: writer, request
func (_m *HttpController) ExportRecordedData(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// ImportRecordedData provides a mock function with given fields: writer, request
func (_m *HttpController) ImportRecordedData(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// RecordingStatus provides a mock function with given fields: writer, request
func (_m *HttpController) RecordingStatus(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// ReplayStatus provides a mock function with given fields: writer, request
func (_m *HttpController) ReplayStatus(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// StartRecording provides a mock function with given fields: writer, request
func (_m *HttpController) StartRecording(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

// StartReplay provides a mock function with given fields: writer, request
func (_m *HttpController) StartReplay(writer http.ResponseWriter, request *http.Request) {
	_m.Called(writer, request)
}

type mockConstructorTestingTNewHttpController interface {
	mock.TestingT
	Cleanup(func())
}

// NewHttpController creates a new instance of HttpController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHttpController(t mockConstructorTestingTNewHttpController) *HttpController {
	mock := &HttpController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
