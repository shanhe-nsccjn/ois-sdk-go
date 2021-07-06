// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package main

import (
	"errors"

	"github.com/DATA-DOG/godog"

	qs "github.com/shanhe-nsccjn/ois-sdk-go/v4/service"
)

// ServiceFeatureContext provides feature context for service.
func ServiceFeatureContext(s *godog.Suite) {
	s.Step(`^initialize OIS service$`, initializeOISService)
	s.Step(`^the OIS service is initialized$`, theOISServiceIsInitialized)

	s.Step(`^list buckets$`, listBuckets)
	s.Step(`^list buckets status code is (\d+)$`, listBucketsStatusCodeIs)
}

// --------------------------------------------------------------------------

func initializeOISService() error {
	return nil
}

func theOISServiceIsInitialized() error {
	if qsService == nil {
		return errors.New("OIS service is not initialized")
	}
	return nil
}

// --------------------------------------------------------------------------

var listBucketsOutput *qs.ListBucketsOutput

func listBuckets() error {
	listBucketsOutput, err = qsService.ListBuckets(nil)
	return err
}

func listBucketsStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(listBucketsOutput.StatusCode), statusCode)
}
