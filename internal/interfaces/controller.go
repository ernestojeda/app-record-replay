// Copyright (c) 2023 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package interfaces

import (
	"net/http"
)

// HttpController defines the interface for implementations that handle the HTTP request
type HttpController interface {
	// AddRoutes adds all the REST routes to the HTTP router
	AddRoutes()
	// StartRecording starts a recording session based on the values in the request.
	// An error is returned if the request data is incomplete or a record or replay session is currently running.
	StartRecording(writer http.ResponseWriter, request *http.Request)
	// CancelRecording cancels the current recording session
	CancelRecording(writer http.ResponseWriter, request *http.Request)
	// RecordingStatus returns the status of the current recording session
	RecordingStatus(writer http.ResponseWriter, request *http.Request)
	// StartReplay starts a replay session based on the values in the request
	// An error is returned if the request data is incomplete or a record or replay session is currently running.
	StartReplay(writer http.ResponseWriter, request *http.Request)
	// CancelReplay cancels the current replay session
	CancelReplay(writer http.ResponseWriter, request *http.Request)
	// ReplayStatus returns the status of the current replay session
	ReplayStatus(writer http.ResponseWriter, request *http.Request)
	// ExportRecordedData returns the data for the last record session
	// An error is returned if the no record session was run or a record session is currently running
	ExportRecordedData(writer http.ResponseWriter, request *http.Request)
	// ImportRecordedData imports data from a previously exported record session.
	// An error is returned if a record or replay session is currently running or the data is incomplete
	ImportRecordedData(writer http.ResponseWriter, request *http.Request)
}
