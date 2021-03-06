// Copyright 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package registry

import mock "github.com/stretchr/testify/mock"

// MockMerger is an autogenerated mock type for the Merger type
type MockMerger struct {
	mock.Mock
}

// Merge provides a mock function with given fields: _a0, _a1
func (_m *MockMerger) Merge(_a0 Entry, _a1 Entry) (Entry, error) {
	ret := _m.Called(_a0, _a1)

	var r0 Entry
	if rf, ok := ret.Get(0).(func(Entry, Entry) Entry); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Entry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Entry, Entry) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ Merger = (*MockMerger)(nil)
