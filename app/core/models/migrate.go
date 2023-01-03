package models

import "worlder-test/internal/config"

func Migrate() {
	config.GetInstanceDb().AutoMigrate(
		&Sensors{}, //migrate sensors
	)
}
