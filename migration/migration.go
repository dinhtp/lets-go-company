package migration

import (
    "gorm.io/gorm"

    "github.com/go-gormigrate/gormigrate/v2"

    "github.com/dinhtp/lets-go-company/migration/versions"
)

func Migrate(db *gorm.DB) error {

    m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID:      "20211002000000",
            Migrate: versions.Version20211002000000,
        },
        {
            ID:      "20211122175000",
            Migrate: versions.Version20211122175000,
        },
    })

    return m.Migrate()
}
