package store

import (
	"sync"

	log "github.com/Sirupsen/logrus"
)

// Employee holds employee info which is stored in the repo
type Employee struct {
	Name           string
	DepartmentName string
	Age            int
}

// Node holds employees at this level and child departments
type node struct {
	name      string
	employees map[string]*Employee
	childDeps map[string]*node
	parent    *node
	// lock
}

// Repo has the root node of the tree
type Repo struct {
	root *node
	lock sync.RWMutex
}

// NewRepo returns a new initialized repository
func NewRepo() *Repo {
	return &Repo{
		root: &node{
			name:      "root",
			employees: map[string]*Employee{},
			childDeps: map[string]*node{},
		},
	}
}

// CreateDepartment creates a new department, return silently if already present
// If parentName is empty, assumes parent as root of tree
func (r *Repo) CreateDepartment(name string, parentName string) bool {
	r.lock.Lock()
	defer r.lock.Unlock()
	var parentNode *node
	//if parentName is empty, assume root
	if parentName == "" {
		parentNode = r.root
	} else {
		parentNode = findNode(r.root, parentName)
	}
	if parentNode == nil {
		log.Error("no parent found while creating department")
		return false
	}
	if parentNode.childDeps[name] != nil {
		return true
	}
	parentNode.childDeps[name] = &node{name: name,
		employees: map[string]*Employee{},
		childDeps: map[string]*node{},
		parent:    parentNode,
	}
	return true
}

// CreateEmployee creates a new employee, return silently if already present
// If departmentName is empty, assumes root
func (r *Repo) CreateEmployee(name string, departmentName string) bool {
	r.lock.Lock()
	defer r.lock.Unlock()
	var departmentNode *node
	if departmentName == "" {
		departmentNode = r.root
		departmentName = "root"
	} else {
		departmentNode = findNode(r.root, departmentName)
	}
	if departmentNode == nil {
		log.Error("no department found while creating employee")
		return false
	}
	departmentNode.employees[name] = &Employee{
		Name:           name,
		DepartmentName: departmentName,
	}
	return true
}

// FindEmployees returns employees under the given root department and with given name
func (r *Repo) FindEmployees(name string, rootDepartmentName string) []Employee {
	r.lock.RLock()
	defer r.lock.RUnlock()
	var departmentNode *node
	if rootDepartmentName == "root" {
		departmentNode = r.root
	} else {
		departmentNode = findNode(r.root, rootDepartmentName)
	}
	if departmentNode == nil {
		return nil
	}
	return findEmployeesByName(departmentNode, name)
}

func findEmployeesByName(node *node, name string) []Employee {
	if node == nil {
		return nil
	}
	employees := []Employee{}
	if node.employees[name] != nil {
		employees = append(employees, *node.employees[name])
	}
	for _, departmentNode := range node.childDeps {
		employees = append(employees, findEmployeesByName(departmentNode, name)...)
	}
	return employees
}

func findNode(root *node, name string) *node {
	if root == nil {
		return nil
	}
	if root.childDeps[name] != nil {
		return root.childDeps[name]
	}
	for _, childDep := range root.childDeps {
		currNode := findNode(childDep, name)
		if currNode != nil {
			return currNode
		}
	}
	return nil
}
