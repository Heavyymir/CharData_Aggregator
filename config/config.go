package config

import (
	"database/sql"
	"github.com/Heavyymir/CharData_Aggregator/catalog"
	"github.com/Heavyymir/CharData_Aggregator/internal/api"
	
)

type Config struct {
	CharDataClient 	*api.Client
	DB				*sql.DB
	Wiki           	catalog.Wiki
	Game           	catalog.Game
	CharacterPage  	string
}
