## ems

Employee management system

## Api's


* Create Department  
If parent_name is empty, assumes root  

      curl -d '{"name": "subadmin", "parent_name": "admin"}' "http://localhost:8080/departments"

* Create Employee  
If department_name is empty, assumes root

      curl -d '{"name": "kiran", "department_name": "admin"}' "http://localhost:8080/employees"  

* Get Employees  
For getting employees under root, use department_name as root  

      curl -d "http://localhost:8080/employees/{empname}/{department_name}"      

## How to run app directly
There is an executable in the repository.Clone the repo and run the executable.

    ./ems

## How to run app from source
1. Clone this app and ``` cd ``` into it.
2. Ensure dependencies ``` dep init and dep ensure  ```
3. Build :
    * ``` go build . ```
    * After this step, executable ems is generated in same directory
4. Run :
    * ``` ./ems ```    

## Test

    curl -d '{"name": "admin", "parent_name": ""}' "http://localhost:8080/departments"
    curl -d '{"name": "subadmin", "parent_name": "admin"}' "http://localhost:8080/departments"
    curl -d '{"name": "subadmin1", "parent_name": "admin"}' "http://localhost:8080/departments"

Create Employee  

    curl -d '{"name": "kiran", "department_name": "admin"}' "http://localhost:8080/employees"
    curl -d '{"name": "kiran", "department_name": "subadmin"}' "http://localhost:8080/employees"  

Get Employees  

    curl -d "http://localhost:8080/employees/kiran/admin"
    {"employee":[{"name":"kiran","department_name":"admin"},{"name":"kiran","department_name":"subadmin"}]}

## Assumptions
1. More domain specific handling or use cases are not covered
2. Employee creation can be done only with a pre-exsiting department

## Further Improvements
1. Improve locking code by using channels and verify race conditions
2. Write tests for all packages and comprehensively.
3. Return more meaningful errors
4. The tree code is not clean.Can be segregated more
5. Model propagation between packages can be cleaned up.
