package models

import (
	"errors"

	"github.com/crackerplace/ems/store"
)

var errMissingDepartmentName = errors.New("department name is empty")
var errCreateDepartment = errors.New("error while creating department")

// Department refers to a department and has a parent
type Department struct {
	Name       string      `json:"name"`
	ParentName string      `json:"parent_name"`
	Repo       *store.Repo `json:"-"`
}

// Create creates a new department if not present
// Returns an error when parent department is not present
func (d *Department) Create() error {
	if d.Name == "" || len(d.Name) == 0 {
		return errMissingDepartmentName
	}
	if !d.Repo.CreateDepartment(d.Name, d.ParentName) {
		return errCreateDepartment
	}
	return nil
}
