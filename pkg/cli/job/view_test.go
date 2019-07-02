/*
Copyright 2019 The Volcano Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package job

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"net/http"
	"net/http/httptest"
	"testing"

	v1alpha1 "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
)

func TestViewJob(t *testing.T) {
	response := v1alpha1.Job{}
	response.Name = "testJob"

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		val, err := json.Marshal(response)
		if err == nil {
			w.Write(val)
		}

	})

	server := httptest.NewServer(handler)
	defer server.Close()

	viewJobFlags.Master = server.URL
	viewJobFlags.Namespace = "test"
	viewJobFlags.JobName = "testJob"

	testCases := []struct {
		Name        string
		ExpectValue error
	}{
		{
			Name:        "viewJob",
			ExpectValue: nil,
		},
	}

	for i, testcase := range testCases {
		err := ViewJob()
		if err != nil {
			t.Errorf("case %d (%s): expected: %v, got %v ", i, testcase.Name, testcase.ExpectValue, err)
		}
	}

}

func TestInitViewFlags(t *testing.T) {
	var cmd cobra.Command
	InitViewFlags(&cmd)

	if cmd.Flag("namespace") == nil {
		t.Errorf("Could not find the flag namespace")
	}
	if cmd.Flag("name") == nil {
		t.Errorf("Could not find the flag name")
	}

}
