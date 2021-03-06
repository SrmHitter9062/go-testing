package main

import (
	"fmt"

	"github.com/SrmHitter9062/go-testing/employee"
)

func main() {
	//get employee
	eData := employee.Emp{
		ID: 1,
	}
	employee := eData.GetEmployee(eData.ID)
	eData.EDetails = employee
	// send notification
	err := eData.Notify()
	if err != nil {
		fmt.Println("Employee could notified, error:", err)
	} else {
		fmt.Println("Employee could not notified")
	}
}
