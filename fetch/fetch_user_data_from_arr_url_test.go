package fetch

import (
	"encoding/json"
	"errors"
	"io"
	"jagaat-technical-task/config"
	"net/http"
	"testing"
)

const respMock = `[
    {
        "_id": "64d39b0582ec3cff5fc7f24e",
        "index": 0,
        "guid": "03ee84da-5a54-493f-8438-60bad7ab6e2a",
        "isActive": true,
        "balance": "$2,633.92",
        "tags": [
            "pariatur",
            "qui",
            "ea",
            "culpa",
            "laboris",
            "laboris",
            "minim"
        ],
        "friends": [
            {
                "id": 0,
                "name": "Koch Valdez"
            },
            {
                "id": 1,
                "name": "Kramer Bush"
            },
            {
                "id": 2,
                "name": "Townsend Church"
            }
        ]
    }
]`

type TestStruct struct {
	name             string
	mockGetResp      *http.Response
	mockGetErr       error
	mockReadAllResp  []byte
	mockReadAllErr   error
	mockUnmarshalErr error
}

func TestFetchAndSaveUserLogicImpl_FetchUserDataFromURLArr(t *testing.T) {
	tests := []TestStruct{
		{
			name: "happy case",
			mockGetResp: &http.Response{
				StatusCode: http.StatusOK,
			},
			mockReadAllResp: []byte(respMock),
		},
		{
			name:       "error fetch api",
			mockGetErr: errors.New("error fetch api"),
		},
		{
			name: "error status not okay",
			mockGetResp: &http.Response{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			name: "error read all",
			mockGetResp: &http.Response{
				StatusCode: http.StatusOK,
			},
			mockReadAllResp: []byte(respMock),
			mockReadAllErr:  errors.New("error read io"),
		},
		{
			name: "error unmarshall",
			mockGetResp: &http.Response{
				StatusCode: http.StatusOK,
			},
			mockReadAllResp:  []byte(respMock),
			mockUnmarshalErr: errors.New("error unmarshall"),
		},
	}

	for _, test := range tests {
		mockDependency(test)
		t.Run(test.name, func(t *testing.T) {
			cfg := config.Config{BaseURL: []string{"http://www.facebook.com"}}
			impl := Impl{}
			impl.FetchUserDataFromURLArr(cfg)
		})
		resetMock()
	}

}

func mockDependency(test TestStruct) {
	Get = func(url string) (resp *http.Response, err error) {
		return test.mockGetResp, test.mockGetErr
	}

	ReadAll = func(r io.Reader) ([]byte, error) {
		return test.mockReadAllResp, test.mockReadAllErr
	}

	Unmarshal = func(data []byte, v any) error {
		return test.mockUnmarshalErr
	}
}

func resetMock() {
	Get = http.Get
	ReadAll = io.ReadAll
	Unmarshal = json.Unmarshal
}
