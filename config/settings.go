package config

type Settings struct {
	db Database
}

func GetSettings() Settings {

	settings := Settings{
		db: GetDatabaseSettings(),
	}

	return settings
}
