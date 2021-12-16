package seeder

import (
    "gorm.io/gorm"
    "time"

    "github.com/bxcodec/faker/v3"
)

type Company struct {
    ID        uint   `faker:"-"`
    Name      string `faker:"domain_name"`
    Phone     string `faker:"toll_free_number"`
    Email     string `faker:"email"`
    Address   string `faker:"sentence"`
    TaxNumber string `faker:"phone_number"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt
}

func FakeCompany() (Company, error) {
    var company Company

    err := faker.FakeData(&company)

    company.CreatedAt = time.Now()
    company.UpdatedAt = time.Now()
    company.DeletedAt = gorm.DeletedAt{}

    return company, err

}
