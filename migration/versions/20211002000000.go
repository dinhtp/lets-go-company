package versions

import (
	"gorm.io/gorm"
	"time"
)

func Version20211002000000(tx *gorm.DB) error {
	type Company struct {
		ID        uint   `gorm:"TYPE:BIGINT(20) UNSIGNED AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
		Name      string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Phone     string `gorm:"TYPE:VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Email     string `gorm:"TYPE:VARCHAR(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
		Address   string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		TaxNumber string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	return tx.AutoMigrate(&Company{})
}
