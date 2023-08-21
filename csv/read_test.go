package csv

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type TestStruct2 struct {
	name        string
	mockStatErr error
	isError     bool
}

func TestLogicImpl_Read(t *testing.T) {
	tests := []TestStruct2{
		{
			name:        "error stat",
			mockStatErr: errors.New("error getting current dir"),
			isError:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockError2(test)
			impl := LogicImpl{}
			_, err := impl.Read()
			if test.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			resetMock()
		})
	}
}

func mockError2(test TestStruct2) {
	Stat = func(name string) (os.FileInfo, error) {
		return nil, test.mockStatErr
	}
}

func resetMock2() {
	Stat = os.Stat
}
