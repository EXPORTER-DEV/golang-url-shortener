package entity

import (
	"testing"
)

type ShortLinkTestCase struct {
	Data     ShortLinkTestData
	IsError  bool
	Expected ShortLinkDatabase
}

type ShortLinkTestData struct {
	JSON string
	ID   string
}

var ShortLinkTestCases = []ShortLinkTestCase{
	ShortLinkTestCase{
		Data: ShortLinkTestData{
			"{\"value\":\"test\"}",
			"TEST",
		},
		IsError: false,
		Expected: ShortLinkDatabase{
			Value:   "test",
			Created: 0,
		},
	},
}

func TestShortLinkDatabaseCreations(t *testing.T) {
	for _, testCase := range ShortLinkTestCases {
		res, err := NewShortLinkDatabase(testCase.Data.JSON)

		if !testCase.IsError && err != nil {
			t.Log("Got unexpected error", err)
			t.Fail()
			continue
		}

		if testCase.IsError && err != nil {
			continue
		}

		if res.Created != testCase.Expected.Created || res.Value != testCase.Expected.Value {
			t.Log("Got mismatch between expected and created structs", "expected:", testCase.Expected, "actual:", res)
			t.Fail()
			continue
		}
	}
}
