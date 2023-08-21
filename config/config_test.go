package config

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestInitializeConfig(t *testing.T) {
	tests := []struct {
		name                 string
		mockGetwdErr         error
		mockOpenFileErr      error
		mockIOReadErr        error
		mockJsonUnarshallErr error
		isError              bool
	}{
		{
			name: "happy case",
		},
		{
			name:         "error case Getwd err",
			mockGetwdErr: errors.New("error getting current directory"),
			isError:      true,
		},
		{
			name:            "error case Open file err",
			mockOpenFileErr: errors.New("error opening file config"),
			isError:         true,
		},
		{
			name:          "error case io Read err",
			mockIOReadErr: errors.New("error reading data from io"),
			isError:       true,
		},
		{
			name:                 "error case Unmarshall err",
			mockJsonUnarshallErr: errors.New("error unmarshall data"),
			isError:              true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Getwd = func() (dir string, err error) {
				return "", test.mockGetwdErr
			}

			Open = func(name string) (*os.File, error) {
				return nil, test.mockOpenFileErr
			}

			ReadAll = func(r io.Reader) ([]byte, error) {
				return []byte{}, test.mockIOReadErr
			}

			Unmarshall = func(data []byte, v any) error {
				return test.mockJsonUnarshallErr
			}

			err := InitializeConfig()
			if test.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
	resetMock()
}

func resetMock() {
	Getwd = os.Getwd
	Open = os.Open
	ReadAll = io.ReadAll
	Unmarshall = json.Unmarshal
}
