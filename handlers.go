// Package main spwans a server daemon and listens for analyse requests.
package main

import (
	"encoding/json"
	"net/http"

	"github.com/crackerplace/ems/models"
	"github.com/crackerplace/ems/store"
	"github.com/gorilla/mux"
)

func createDepartmentHandler(repo *store.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var dep models.Department
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&dep); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		dep.Repo = repo
		if err := dep.Create(); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer r.Body.Close()
		respondWithJSON(w, http.StatusCreated, dep)
	}
}

func createEmployeeHandler(repo *store.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var emp models.Employee
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&emp); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()
		emp.Repo = repo
		if err := emp.Create(); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, emp)
	}
}

// EmployeesByNameResponse has the employees with a specific name and falling under a root department
type EmployeesByNameResponse struct {
	Employees []models.Employee `json:"employee"`
}

func getEmployeesHandler(repo *store.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		departmentName := mux.Vars(r)["departmentName"]
		emp := models.Employee{}
		emp.Repo = repo
		employeesResp := EmployeesByNameResponse{
			Employees: emp.EmployeesByNameAndDepartment(name, departmentName),
		}
		respondWithJSON(w, http.StatusCreated, employeesResp)
	}
}
