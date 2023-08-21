package read_and_search

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"jagaat-technical-task/csv"
	"strconv"
	"testing"
)

var (
	mockReadCsv = [][]string{
		{"ID", "Index", "GUID", "IsActive", "Balance", "Tags", "Friend"},
		{"64d39b0582ec3cff5fc7f24e", "0", "03ee84da-5a54-493f-8438-60bad7ab6e2a", "1", "$2,633.92", "pariatur|qui|ea|culpa|laboris|laboris|minim", "[{\"id\":0,\"name\":\"Koch Valdez\"},{\"id\":1,\"name\":\"Kramer Bush\"},{\"id\":2,\"name\":\"Townsend Church\"}]"},
		{"64d39b0506da562378c0f321", "1", "bfb84280-21a2-4199-833f-4f54d98fb15b", "1", "$3,626.10", "irure|nostrud|ipsum|consectetur|consectetur|occaecat|consectetur", "[{\"id\":0,\"name\":\"Therese Dorsey\"},{\"id\":1,\"name\":\"Gilliam Stephens\"},{\"id\":2,\"name\":\"Leblanc Odonnell\"}]"},
	}
)

type TestStruct struct {
	name              string
	tagSearch         string
	mockReadCsvResult [][]string
	mockReadCsvErr    error
	isError           bool
	mockAtoi1Err      error
	mockAtoi2Err      error
	mockUnmarshalErr  error
	mockMarshalErr    error
}

func TestReadAndSearchLogicImpl_ReadAndSearch(t *testing.T) {
	tests := []TestStruct{
		{
			name:              "happy case",
			tagSearch:         "pariatur",
			mockReadCsvResult: mockReadCsv,
		},
		{
			name:              "error unmarshall",
			tagSearch:         "pariatur",
			mockReadCsvResult: mockReadCsv,
			mockUnmarshalErr:  errors.New("error unmarshal"),
		},
		{
			name:              "error marshal",
			tagSearch:         "pariatur",
			mockReadCsvResult: mockReadCsv,
			mockMarshalErr:    errors.New("error marshal"),
		},
		{
			name:              "error atoi1 skip loop",
			tagSearch:         "pariatur",
			mockReadCsvResult: mockReadCsv,
			mockAtoi1Err:      errors.New("error converting string to int"),
		},
		{
			name:              "error atoi2 skip loop",
			tagSearch:         "pariatur",
			mockReadCsvResult: mockReadCsv,
			mockAtoi2Err:      errors.New("error converting string to int"),
		},
		{
			name:           "error reading csv",
			mockReadCsvErr: errors.New("error reading csv"),
			isError:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockError(test)
			mockCsv := csv.MockICSV{}
			mockCsv.On("Read", mock.Anything).Return(test.mockReadCsvResult, test.mockReadCsvErr)
			impl := ReadAndSearchLogicImpl{
				CSVLogic: &mockCsv,
			}
			err := impl.ReadAndSearch(test.tagSearch)
			if test.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
		resetMock()
	}
}

func mockError(test TestStruct) {
	Atoi1 = func(s string) (int, error) {
		return 0, test.mockAtoi1Err
	}

	Atoi2 = func(s string) (int, error) {
		return 1, test.mockAtoi2Err
	}

	Unmarshal = func(data []byte, v any) error {
		return test.mockUnmarshalErr
	}

	if test.mockMarshalErr != nil {
		Marshal = func(v any) ([]byte, error) {
			return nil, test.mockMarshalErr
		}
	}
}

func resetMock() {
	Atoi1 = strconv.Atoi
	Atoi2 = strconv.Atoi
	Unmarshal = json.Unmarshal
	Marshal = json.Marshal
}
