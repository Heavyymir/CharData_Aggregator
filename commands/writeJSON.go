package commands

import(
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Heavyymir/CharData_Aggregator/internal/parsers/bbcf"
)


//The below helper is designed to create a Json output for the program to store outputs
func writeMovesJSON(character string, moves []bbcf.Move) error {
	filename := safeFileName(character) + ".json"

	data, err := json.MarshalIndent(moves, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return err
	}

	fmt.Printf("Wrote  %s\n", filename)
	return nil
}


func safeFileName(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "_"))
}
