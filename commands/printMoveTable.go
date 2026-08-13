package commands

import (
	"fmt"

	"github.com/Heavyymir/CharData_Aggregator/internal/models"
)

// helper to print parsed CharData to console in a readable format
func printMoveTable(moves []models.Move) {
	for _, move := range moves {
		fmt.Printf("\n%s %s\n", move.Name, move.Input)
		fmt.Println("---------------------------------------------")

		for _, header := range move.Headers {
			cell := move.FrameData[header]

			if cell.Tooltip != "" {
				fmt.Printf("%-18s %-12s (%s)\n",
					header, cell.Value, cell.Tooltip)
			} else {
				fmt.Printf("%-18s %s\n", header, cell.Value)
			}
		}

		for _, note := range move.Notes {
			fmt.Printf("Note: %s\n", note)
		}

		if move.Description != "" {
			fmt.Printf("Description: %s\n", move.Description)
		}
	}
}
