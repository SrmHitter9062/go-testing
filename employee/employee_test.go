package employee

import (
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)


type mockObj struct {
	mock.Mock
}

func (m *mockObj) Publish(data []byte) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestNotify(t *testing.T) {
	mock := new(mockObj)
	emp := Emp{
		ID:        1,
		EDetails:   EDetails{
			ID:      1,
			Name:    "Bill",
			Address: "54, X colony, mumbai, 341232",
			Salary: SDetails{
				Sent:   true,
				Month:  "February",
				Amount: 100000,
				Bank:   "SBI",
			},
		},
	}
	emp.Publisher = mock
	mock.On("Publish", emp.EDetails).Return(nil)
	res := emp.Notify()
	assert.Equal(t,nil,res)
}

