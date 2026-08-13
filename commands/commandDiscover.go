package commands

import (
	"fmt"
	"strings"

	"github.com/Heavyymir/CharData_Aggregator/config"
	"github.com/Heavyymir/CharData_Aggregator/internal/CharDataCache"
	"github.com/Heavyymir/CharData_Aggregator/internal/discovery"
)

func commandDiscover(cfg *config.Config, args ...string) error {
	if cfg.Wiki.Name == "" || cfg.Game.Name == "" {
		return fmt.Errorf("Select a wiki and game first")
	}

	wikiSlug := strings.ToUpper(cfg.Game.Slug)

	indexURL := cfg.Wiki.URL + wikiSlug

	data, err := cfg.CharDataClient.Fetch(indexURL)
	if err != nil {
		return err
	}

	characters, err := discovery.DiscoveredChars(data, wikiSlug)
	if err != nil {
		return err
	}

	for _, character := range characters {
		fmt.Printf("%s -> %s\n", character.Name, character.URL)
	}

	filename := cfg.Game.Slug + "_characters" + ".json"

	if err := CharDataCache.SaveCharacters(
		filename,
		cfg.Game.Slug,
		characters,
	); err != nil {
		return err
	}

	cache, err := discovery.LoadCharCache(filename)
	if err != nil {
		return err
	}


	fmt.Printf("Discovered %d characters for %s\n", len(characters), cfg.Game.Name)
	fmt.Printf("Loaded %d characters\n", len(cache.Characters))
	return nil
}
