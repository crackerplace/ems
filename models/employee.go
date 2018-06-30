package models

import (
	"errors"

	"github.com/crackerplace/ems/store"
)

var errMissingName = errors.New("name is empty")
var errCreateEmployee = errors.New("error while creating employee")

// Employee holds employee info
type Employee struct {
	Name           string `json:"name"`
	DepartmentName string `json:"department_name"`
	//Age            int         `json:"age"`
	Repo *store.Repo `json:"-"`
}

// Create creates a new employee if not present
// If present, returns silently
// If departmentName is not present, returns an error
func (e *Employee) Create() error {
	if isEmpty(e.Name) {
		return errMissingName
	}
	if !e.Repo.CreateEmployee(e.Name, e.DepartmentName) {
		return errCreateEmployee
	}
	return nil
}

// EmployeesByNameAndDepartment returns employees with given name and with root department as given
func (e *Employee) EmployeesByNameAndDepartment(name, rootDepartmentName string) []Employee {
	employees := []Employee{}
	if e.Repo.FindEmployees(name, rootDepartmentName) != nil {
		for _, emp := range e.Repo.FindEmployees(name, rootDepartmentName) {
			employees = append(employees, Employee{Name: emp.Name, DepartmentName: emp.DepartmentName})
		}
	}
	return employees
}

func isEmpty(val string) bool {
	if val == "" || len(val) == 0 {
		return true
	}
	return false
}
