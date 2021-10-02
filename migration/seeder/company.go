package seeder

import (
    "gorm.io/gorm"
    "time"

    "github.com/bxcodec/faker/v3"
)

type Company struct {
    gorm.Model
    Name      string `faker:"domain_name"`
    Phone     string `faker:"phone_number"`
    Email     string `faker:"email"`
    Address   string `faker:"sentence"`
    TaxNumber string `faker:"toll_free_number"`
}

func FakeCompany() (Company, error) {
    var company Company

    err := faker.FakeData(&company)

    company.CreatedAt = time.Now()
    company.UpdatedAt = time.Now()

    return company, err

}
