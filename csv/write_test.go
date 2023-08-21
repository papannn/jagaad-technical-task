package csv

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"jagaat-technical-task/dto"
	"os"
	"testing"
)

type TestStruct struct {
	name            string
	userArr         []dto.User
	mockCreateErr   error
	mockWriteAllErr error
	mockMarshalErr  error
	isError         bool
}

func TestLogicImpl_Write(t *testing.T) {
	tests := []TestStruct{
		{
			name: "happy case",
			userArr: []dto.User{
				{ID: "1", Index: 0, GUID: "guid", IsActive: true, Balance: "balance", Tags: []string{"tag"}, Friends: []dto.Friend{}},
			},
		},
		{
			name:          "error create csv",
			mockCreateErr: errors.New("error create csv"),
			isError:       true,
		},
		{
			name: "error marshal data",
			userArr: []dto.User{
				{ID: "1", Index: 0, GUID: "guid", IsActive: true, Balance: "balance", Tags: []string{"tag"}, Friends: []dto.Friend{
					{
						ID:   0,
						Name: "Taufan",
					},
				}},
			},
			mockMarshalErr: errors.New("error marshal data"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockError(test)
			impl := &LogicImpl{}
			err := impl.Write(test.userArr)
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
	if test.mockCreateErr != nil {
		Create = func(name string) (*os.File, error) {
			return &os.File{}, test.mockCreateErr
		}
	}

	Marshal = func(v any) ([]byte, error) {
		return nil, test.mockMarshalErr
	}
}

func resetMock() {
	Create = os.Create
	NewWriter = csv.NewWriter
	Marshal = json.Marshal
}
