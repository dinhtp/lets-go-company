package migration

import (
	"gorm.io/gorm"

    "github.com/go-gormigrate/gormigrate/v2"

	"github.com/dinhtp/lets-go-company/migration/versions"
)

func Migrate(db *gorm.DB) error {
    m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID:      "20211215000000",
            Migrate: versions.Version20211215000000,
        },
    })

    return m.Migrate()
}