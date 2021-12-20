# Let's GO Company

## What to do
In this service, you will develop a backend server side functions that handle the **company** and **employee** API
resources.This service will contain 2 sub-services which are the Rest and gRPC service.
Rest service will handle the http requests while gRPC will handle the rpc request to the API resources.

The company and employee ERD can be referred below:
![company-employee ERD](./asset/company-employee.png)

## API Output
### Company API Output
#### Get a company by ID.
    - URL: [GET] {company_url}/go/company/{id}
    - Response: 
        {
            "id": "string",
            "name": "string",
            "phone": "string",
            "email": "string",
            "address": "string",
            "tax_number": "string",
            "total_employee": 0, // calculate the total employee
            "total_project": 0, // calculate the total "active" project
            "created_at": "string",
            "updated_at": "string"
        }
#### Create a company.
    - URL: [POST] {company_url}/go/company
    - Payload:
        {
            "name": "string",
            "phone": "string",
            "email": "string",
            "address": "string",
            "tax_number": "string"
        }
    - Response:
        {
            "id": "string",
            "name": "string",
            "phone": "string",
            "email": "string",
            "address": "string",
            "tax_number": "string",
            "total_employee": 0, // calculate the total employee
            "total_project": 0, // calculate the total "active" project
            "created_at": "string",
            "updated_at": "string"
        }
#### Update a company by ID.
    - URL: [PUT] {company_url}/go/company/{id}
    - Payload:
        {
            "id": "string",
            "name": "string",
            "phone": "string",
            "email": "string",
            "address": "string",
            "tax_number": "string"
        }
    - Response:
        {
            "id": "string",
            "name": "string",
            "phone": "string",
            "email": "string",
            "address": "string",
            "tax_number": "string",
            "total_employee": 0, // calculate the total employee
            "total_project": 0, // calculate the total "active" project
            "created_at": "string",
            "updated_at": "string"
        }
#### Delete a company by ID.
    - URL: [DELETE] {company_url}/go/company/{id}
    - Status: 200
#### List a company by page, limit and filter by "name", "phone", "email"
    - URL: [GET] {company_url}/go/compannies
    - Query: ?page=0&limit=0&search_value=string&search_fields=name,phone,email
    - Response:
        {
            "items": [
                {
                    "id": "string",
                    "name": "string",
                    "phone": "string",
                    "email": "string",
                    "address": "string",
                    "tax_number": "string",
                    "created_at": "string",
                    "updated_at": "string"
                },
                ...
            ]
            "max_page": 0,
            "total_count": 0,
            "page": 0,
            "limit": 0,
        }

### Employee API Output
#### Get an employee by ID.
    - URL: [GET] {company_url}/go/employee/{id}
    - Response:
        {
            "id": "string",
            "company_id": "string",
            "name": "string",
            "email": "string",
            "dob": "string",
            "gender": "string",
            "role": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Create an employee for a specific company.
    - URL: [POST] {company_url}/go/company/{company_id}/employee
    - Payload:
        {
            "company_id": "string",
            "name": "string",
            "email": "string",
            "dob": "string",
            "gender": "string", // allowed values: "male", "female", "not_specified"
            "role": "string", // allowed values: "manager", "staff", "intern"
        }
    - Response:
        {
            "id": "string",
            "company_id": "string",
            "name": "string",
            "email": "string",
            "dob": "string",
            "gender": "string",
            "role": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Update an employee by ID.
    - URL: [POST] {company_url}/go/employee/{id}
    - Payload:
        {
            "id": "string",
            "company_id": "string",
            "name": "string",
            "email": "string",
            "dob": "string",
            "gender": "string", // allowed values: "male", "female", "not_specified"
            "role": "string", // allowed values: "manager", "staff", "intern"
        }
    - Response:
        {
            "id": "string",
            "company_id": "string",
            "name": "string",
            "email": "string",
            "dob": "string",
            "gender": "string",
            "role": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Delete an employee by ID.
    - URL: [DELETE] {company_url}/go/employee/{id}
    - Status: 200
#### List employee by company id, page, limit and filter by "name", "email"
    - URL: [GET] {company_url}/go/company/{company_id}/employees
    - Query: ?page=0&limit=0&search_value=string&search_fields=name,email
    - Response:
        {
            "items": [
                {
                    "id": "string",
                    "company_id": "string",
                    "name": "string",
                    "email": "string",
                    "dob": "string",
                    "gender": "string",
                    "role": "string",
                    "created_at": "string",
                    "updated_at": "string"
                },
                ...
            ]
            "max_page": 0,
            "total_count": 0,
            "page": 0,
            "limit": 0,
        }

> NOTE: DO NOT commit changes directly into the master branch. Make your own master branch.