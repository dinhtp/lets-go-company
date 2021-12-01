package seeder

import (
    "time"

    "gorm.io/gorm"
    "github.com/bxcodec/faker/v3"
)

type Employee struct {
    gorm.Model
    Name      string `faker:"domain_name"`
    Phone     string `faker:"toll_free_number"`
    Email     string `faker:"email"`
    Address   string `faker:"sentence"`
    TaxNumber string `faker:"phone_number"`
}

func FakeEmployee() (Employee, error) {
    var employee Employee

    err := faker.FakeData(&employee)

    employee.CreatedAt = time.Now()
    employee.UpdatedAt = time.Now()
    employee.DeletedAt = gorm.DeletedAt{}

    return employee, err

}
