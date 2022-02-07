package seeder

import "gorm.io/gorm"

type Seeder interface {
	TableName() string
	Seed(*gorm.DB) error
}

func RunSeeders(db *gorm.DB) {
	seeders := []Seeder{
		Users,
	}
	for _, seeder := range seeders {
		if err := seeder.Seed(db); err != nil {
			panic("seeder failed: " + seeder.TableName())
		}
	}
}
