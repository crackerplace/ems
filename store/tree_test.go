package store

import "testing"

func TestEmptyRepo(t *testing.T) {
	repo := NewRepo()
	employees := repo.FindEmployees("kiran", "admin")
	if len(employees) != 0 {
		t.Errorf("Expected no employees. Got employees %v", employees)
	}
}

func TestCreateEmployeeWithoutValidDepartment(t *testing.T) {
	repo := NewRepo()
	if repo.CreateEmployee("kiran", "admin") {
		t.Errorf("Expected employee creation to fail. But employee creation succesful")
	}
}

func TestCreateEmployeeWithEmptyName(t *testing.T) {
	repo := NewRepo()
	if repo.CreateEmployee("", "admin") {
		t.Errorf("Expected employee creation to fail. But employee creation succesful")
	}
}

func TestCreateEmployeeUnderValidDepartment(t *testing.T) {
	repo := NewRepo()
	repo.CreateDepartment("admin", "")
	repo.CreateEmployee("kiran", "admin")
	if len(repo.FindEmployees("kiran", "admin")) != 1 {
		t.Errorf("Expected employee creation to suceed. But employee creation failed")
	}
}

func TestCreateEmployeeAtRoot(t *testing.T) {
	repo := NewRepo()
	repo.CreateEmployee("kiran", "")
	if len(repo.FindEmployees("kiran", "root")) != 1 {
		t.Errorf("Expected employee creation to suceed. But employee creation failed")
	}
}

// expose a method to find employee at specific department node in repo to test creation
// func TestCreateEmployeeAtSpecifcDepartmentNode(t *testing.T) {
// 	repo := NewRepo()
// 	repo.CreateDepartment("admin", "")
// 	repo.CreateEmployee("kiran", "admin")
// 	if len(repo.FindEmployees("kiran", "admin")) != 1 {
// 		t.Errorf("Expected employee creation to fail. But employee creation succesful")
// 	}
// }

func TestCreateDepartmentWithoutValidParent(t *testing.T) {
	repo := NewRepo()
	if repo.CreateEmployee("kiran", "admin") {
		t.Errorf("Expected department creation to fail. But department creation succesful")
	}
}

//
// func TestCreateDepartment(t *testing.T) {
// }
//
// func TestCreateEmployee(t *testing.T) {
// }
