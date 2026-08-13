package catalog

import "sort"

func WikiSuggestions() []string {
	var suggestions []string
	for _, wiki := range Wikis {
		suggestions = append(suggestions, wiki.Name) 
	}
	sort.Strings(suggestions)
	return suggestions
}


func GameSuggestions() []string {
	var suggestions []string
	for _, wiki := range Wikis {
		for _, game := range wiki.Games {
			suggestions = append(suggestions, game.Slug)
		}
	}
	sort.Strings(suggestions)
	return suggestions
}
