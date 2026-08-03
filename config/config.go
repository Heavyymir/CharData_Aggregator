package config

import (
	"github.com/Heavyymir/CharData_Aggregator/catalog"
	"github.com/Heavyymir/CharData_Aggregator/internal/api"
)

type Config struct {
	CharDataClient *api.Client
	Wiki           catalog.Wiki
	Game           catalog.Game
	CharacterPage  string
}
