package main

import "testing"

type GetReportUriTestCase struct {
	Name           string
	Uri            string
	ExpectedError  bool
	ExpectedRet    string
	FailureMessage string
}

func Test_GetReportUri(t *testing.T) {
	testCases := []GetReportUriTestCase{
		{
			Name:           "",
			Uri:            "",
			ExpectedError:  true,
			ExpectedRet:    "",
			FailureMessage: "Didn't report error when both and uri are empty",
		},
		{
			Name:           "fake",
			Uri:            "abc",
			ExpectedError:  true,
			ExpectedRet:    "",
			FailureMessage: "Didn't report error when both and uri are specified",
		},
		{
			Name:           "fake111",
			Uri:            "",
			ExpectedError:  true,
			ExpectedRet:    "",
			FailureMessage: "Failed to report error if built-in report doesn't exist",
		},
		{
			Name:           "fake",
			Uri:            "",
			ExpectedError:  false,
			ExpectedRet:    "http://localhost:8080/reports/fake",
			FailureMessage: "Failed to generate built-in report uri based on name",
		},
		{
			Name:           "",
			Uri:            "abc://http//abc",
			ExpectedError:  true,
			ExpectedRet:    "",
			FailureMessage: "Failed to report error when uri scheme is invalid",
		},
		{
			Name:           "",
			Uri:            "http:///abc",
			ExpectedError:  true,
			ExpectedRet:    "",
			FailureMessage: "Failed to report error when uri host is invalid",
		},
		{
			Name:           "",
			Uri:            "http://bing.com/abc",
			ExpectedError:  false,
			ExpectedRet:    "http://bing.com/abc",
			FailureMessage: "Failed to generate external report url based on uri",
		},
	}

	for _, tc := range testCases {
		req := DownloadReq{
			Name: tc.Name,
			Uri:  tc.Uri,
		}

		ret, err := getReportUri(req)
		if (tc.ExpectedError == true && err == nil) ||
			(tc.ExpectedError == false && err != nil) ||
			ret != tc.ExpectedRet {
			t.Error(tc.FailureMessage)
		}
	}
}
