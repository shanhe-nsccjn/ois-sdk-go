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

package signer

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/pengsrc/go-shared/convert"
	"github.com/stretchr/testify/assert"

	"github.com/shanhe-nsccjn/ois-sdk-go/v4/utils"
)

func TestOISSignerWriteSignature(t *testing.T) {
	url := "https://ois.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0&other=abc"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", convert.TimeToString(time.Time{}, convert.RFC822))
	httpRequest.Header.Set("X-QS-Test-2", "Test 2")
	httpRequest.Header.Set("X-QS-Test-1", "Test 1")
	assert.Nil(t, err)

	s := OISSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteSignature(httpRequest)
	assert.Nil(t, err)

	signature := "QS ENV_ACCESS_KEY_ID:bvglZF9iMOv1RaCTxPYWxexmt1UN2m5WKngYnhDEp2c="
	assert.Equal(t, signature, httpRequest.Header.Get("Authorization"))
}

func TestOISSignerWriteSignatureWithXQSDate(t *testing.T) {
	url := "https://ois.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0&other=abc"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", convert.TimeToString(time.Time{}, convert.RFC822))
	httpRequest.Header.Set("X-QS-Date", convert.TimeToString(time.Time{}, convert.RFC822))
	httpRequest.Header.Set("X-QS-Test-2", "Test 2")
	httpRequest.Header.Set("X-QS-Test-1", "Test 1")
	assert.Nil(t, err)

	s := OISSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteSignature(httpRequest)
	assert.Nil(t, err)

	signature := "QS ENV_ACCESS_KEY_ID:qkY+tOMdqfDAVv+ZBtlWeEBxlbyIKaQmj5lQlylENzo="
	assert.Equal(t, signature, httpRequest.Header.Get("Authorization"))
}

func TestOISSignerWriteSignatureChinese(t *testing.T) {
	url := "https://zone.ois.com/bucket-name/??????"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", convert.TimeToString(time.Time{}, convert.RFC822))
	assert.Nil(t, err)

	s := OISSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteSignature(httpRequest)
	assert.Nil(t, err)

	signature := "QS ENV_ACCESS_KEY_ID:XsTXX50kzqBf92zLG1aIUIJmZ0hqIHoaHgkumwnV3fs="
	assert.Equal(t, signature, httpRequest.Header.Get("Authorization"))
}

func TestOISSignerWriteQuerySignature(t *testing.T) {
	url := "https://ois.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", convert.TimeToString(time.Time{}, convert.RFC822))
	httpRequest.Header.Set("X-QS-Test-2", "Test 2")
	httpRequest.Header.Set("X-QS-Test-1", "Test 1")
	assert.Nil(t, err)

	s := OISSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteQuerySignature(httpRequest, 3600)
	assert.Nil(t, err)

	targetURL := "https://ois.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0&access_key_id=ENV_ACCESS_KEY_ID&expires=3600&signature=GRL3p3NOgHR9CQygASvyo344vdnO1hFke6ZvQ5mDVHM="
	assert.Equal(t, httpRequest.URL.String(), targetURL)
}

func TestOISSignerWriteQuerySignatureWithXQSDate(t *testing.T) {
	url := "https://ois.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", convert.TimeToString(time.Time{}, convert.RFC822))
	httpRequest.Header.Set("X-QS-Date", convert.TimeToString(time.Time{}, convert.RFC822))
	httpRequest.Header.Set("X-QS-Test-2", "Test 2")
	httpRequest.Header.Set("X-QS-Test-1", "Test 1")
	assert.Nil(t, err)

	s := OISSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteQuerySignature(httpRequest, 3600)
	assert.Nil(t, err)

	targetURL := "https://ois.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0&access_key_id=ENV_ACCESS_KEY_ID&expires=3600&signature=plFxMFP1EzKVtdF%2BbApT8rhW9AUAIWfmZcOGH3m27t0="
	assert.Equal(t, httpRequest.URL.String(), targetURL)
}

func TestOISSigner_WriteQuerySignatureWithDisposition(t *testing.T) {
	objKey := "??????.jpg"
	dlName := utils.URLQueryEscape(objKey)
	disposition := fmt.Sprintf("attachment; filename=\"%s\"; filename*=utf-8''%s", dlName, dlName)
	// Let's start with a base url
	u, _ := url.Parse("https://ois.com")
	params := u.Query()
	params.Set("response-content-disposition", disposition)
	u.RawQuery = params.Encode()
	httpRequest, _ := http.NewRequest("GET", u.String(), nil)
	httpRequest.Header.Set("Date", convert.TimeToString(time.Time{}, convert.RFC822))

	s := OISSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err := s.WriteQuerySignature(httpRequest, 3600)
	assert.Nil(t, err)
	targetURL := "https://ois.com?response-content-disposition=attachment%3B+filename%3D%22%25E4%25B8%25AD%25E6%2596%2587.jpg%22%3B+filename%2A%3Dutf-8%27%27%25E4%25B8%25AD%25E6%2596%2587.jpg&access_key_id=ENV_ACCESS_KEY_ID&expires=3600&signature=8mZxt4VXwmiiERhfytHyuySjWZD/VPMC3kAy%2BANwHoI="
	assert.Equal(t, httpRequest.URL.String(), targetURL)
}

func TestOISSigner_WriteSignatureWithCname(t *testing.T) {
	url := "https://ois.com/bucket-name?cname"
	httpRequest, err := http.NewRequest("PUT", url, nil)
	httpRequest.Header.Set("Date", convert.TimeToString(time.Time{}, convert.RFC822))
	httpRequest.Header.Set("X-QS-Test-2", "Test 2")
	httpRequest.Header.Set("X-QS-Test-1", "Test 1")
	assert.Nil(t, err)

	s := OISSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteSignature(httpRequest)
	assert.Nil(t, err)

	signature := "QS ENV_ACCESS_KEY_ID:Kqc4/+6T7zfSrUPgkvswHbtL4ESch9vOVQP0nPwlkBs="
	assert.Equal(t, signature, httpRequest.Header.Get("Authorization"))
}

func TestOISSinger_WriteSignatureWithVirtualStyle(t *testing.T) {
	url := "https://bucket-name.ois.com/?acl"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", convert.TimeToString(time.Time{}, convert.RFC822))
	assert.Nil(t, err)

	s := OISSigner{
		AccessKeyID:            "ENV_ACCESS_KEY_ID",
		SecretAccessKey:        "ENV_SECRET_ACCESS_KEY",
		EnableVirtualHostStyle: true,
	}

	err = s.WriteSignature(httpRequest)
	assert.Nil(t, err)

	signature := "QS ENV_ACCESS_KEY_ID:y0rP8LnEGUiV/JcuCMHZC4OMRZNzKRw9w7AnqehJZW8="
	assert.Equal(t, signature, httpRequest.Header.Get("Authorization"))
}
