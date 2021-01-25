package migration

import (
	"fmt"

	log "github.com/Bendomey/RideHail/account/internal/logger"

	"github.com/Bendomey/RideHail/account/internal/orm/migration/jobs"
	"github.com/Bendomey/RideHail/account/internal/orm/models"
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func updateMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Admin{},
		&models.Rider{},
		&models.Customer{},
	)
	return err
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Info("[Migration.InitSchema] Initializing database schema")
		db.Exec("create extension \"uuid-ossp\";")
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		// Add more jobs, etc here
		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedSuperAdmin,
	})
	return m.Migrate()
}
