package seeder

import (
    "gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
    err := CreateFakeCompany(db)
    if nil != err {
        return err
    }

    return nil
}

func CreateFakeCompany(db *gorm.DB) error {
    for i := 1; i <= 10; i++ {
        co, err := FakeCompany()
        if nil != err {
            return err
        }

        err = db.Create(&co).Error
        if nil != err {
            return err
        }
    }
    return nil
}
