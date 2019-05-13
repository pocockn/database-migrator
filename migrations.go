package main

import (
	"database-migrator/migrations"
	gormigrate "gopkg.in/gormigrate.v1"
)

// GenerateMigrations generates migrations for our db.
func GenerateMigrations() []*gormigrate.Migration {
	var gormMigrations []*gormigrate.Migration

	allMigrations := getMigrations()

	for _, migration := range allMigrations {
		gormMigrations = append(gormMigrations, migration())
	}

	return gormMigrations
}

func getMigrations() []func() *gormigrate.Migration {
	return []func() *gormigrate.Migration{
		migrations.ShoutTableMigration,
	}
}
