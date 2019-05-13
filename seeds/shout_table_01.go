package seeds

import (
	"github.com/jinzhu/gorm"
	"github.com/pocockn/models/api/shouts"
)

// SeedShoutTable seeds the team table with test data.
func SeedShoutTable(db *gorm.DB) error {
	teams := seedTeams()

	for _, team := range teams {
		err := db.Create(&team).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedShouts() []shouts.Shout {
	var shouts []shouts.Shout

	shouts = append(
		shouts,
		shouts.Shout{

		},
		shouts.Shout{

		},
	)

	return shouts
}
