package fetch_and_save_user_cmd

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"jagaat-technical-task/csv"
	"jagaat-technical-task/dto"
	"jagaat-technical-task/fetch"
	"testing"
)

func TestFetchAndSaveUserLogicImpl_FetchAndSaveUser(t *testing.T) {
	tests := []struct {
		name              string
		mockFetchUserData []dto.User
		mockWriteCsvErr   error
		isError           bool
	}{
		{
			name: "happy case",
		},
		{
			name:            "error write case",
			mockWriteCsvErr: errors.New("error writing to csv"),
			isError:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockFetch := fetch.MockIFetch{}
			mockCsv := csv.MockICSV{}

			mockFetch.On("FetchUserDataFromURLArr", mock.Anything).Return(test.mockFetchUserData)
			mockCsv.On("Write", mock.Anything).Return(test.mockWriteCsvErr)

			impl := FetchAndSaveUserLogicImpl{
				CSVLogic:   &mockCsv,
				FetchLogic: &mockFetch,
			}

			err := impl.FetchAndSaveUser()
			if test.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
