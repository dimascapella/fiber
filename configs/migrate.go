package configs

import "gorm.io/gorm"

func Migrate(db *gorm.DB, models []interface{}) {
	db.Migrator().DropTable(models...)
	db.AutoMigrate(models...)
}
