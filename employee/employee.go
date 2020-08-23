package employee

import (
	"../publisher"
	"encoding/json"
	"fmt"
)

type Emp struct {
	ID int
	EDetails EDetails
	// embedding dependent interface
	Publisher publisher.PublishInterface
}

type EDetails struct {
	ID int
	Name string
	Address string
	Salary SDetails
}

type SDetails struct {
	Sent bool
	Month string
	Amount int
	Bank string
}

func (e *Emp) GetEmployee(ID int) EDetails {
	// some  code to get employee details
	return EDetails{
		ID:      1,
		Name:    "Bill",
		Address: "54, X colony, mumbai, 341232",
		Salary: SDetails{
			Sent:   true,
			Month:  "February",
			Amount: 100000,
			Bank:   "SBI",
		},
	}
}

func (e *Emp) Notify() error {
	e.Publisher = &publisher.Publisher{}
	// get data
	data, err := json.Marshal(e.EDetails)
	if err != nil {
		fmt.Println("error in encoding the employee details")
		return err
	}
	pErr := e.Publisher.Publish(data)
	return pErr
}
