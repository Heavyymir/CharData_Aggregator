package ggst

import (
        "bytes"
        "strings"
        "fmt"

        "github.com/PuerkitoBio/goquery"
        "github.com/Heavyymir/CharData_Aggregator/internal/models"
)

func GGSTCharPageParser(data []byte) ([]models.Move, error) {

        // Use goquery to create a readable datapoint from page HTML data
        doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
        if err != nil {
                return nil, err
        }

        // Initialise a slice of Move structs
        var moves []models.Move

        // Find all attack containers in the created doc
        doc.Find(".attack-container").Each(func(i int, container *goquery.Selection) {
	
                // Initialise move struct maps/slices for elements
                move := models.Move{
                        FrameData: make(map[string]models.Cell),
                        Headers: []string{},
                }

				// Find move names
				heading := container.PrevAll().Filter("div.mw-heading3").First()
				if heading.Length() > 0 {
					move.Name = strings.TrimSpace(heading.Find("h3").Text())
				}
				
                // Assign move inputs
                inputNode := container.Prev()
                input := strings.TrimSpace(inputNode.Find(".input-badge").Text())
                move.Input = strings.Join(strings.Fields(input), " ")

                // Find all of the framedatagridheaders in the doc
                container.Find(".frameDataGridHeader > div").Each(func(_ int, cell *goquery.Selection) {
                        // Clone the created goquery
                        visible := cell.Clone()

                        // Find the tooltip text inside the copied goquery and remove it
                        visible.Find(".tooltiptext").Remove()

                        // Append the remaining text (the raw header without explanation) to the headers slice
                        header := strings.TrimSpace(visible.Text())
                        move.Headers = append(move.Headers, header)
                        })

                // Find the framedata grid rows present in the Movecontainers. Map data for each row to a header.
                container.Find(".frameDataGridRow > div").Each(func(i int, cell *goquery.Selection) {
                                // if headers is shorter than the number of framedata rows, safetly return
                                if i >= len(move.Headers) {
                                        return
                                }

                                // Parse the created goquery with the helper and assign the header data to move.FrameData
                                move.FrameData[move.Headers[i]] = parseCell(cell)
                        })

                // Find the attack body info containing the move tooltips from the wiki
                container.Find(".attack-info-body > ul > li").Each(func(_ int, note *goquery.Selection) {
                        // Append the text element of the created goquery to the Notes element of the created Move struct
                        noteText := visibleText(note)
                        if noteText != "" {
                        	move.Notes = append(move.Notes, noteText)
                        }
                })

                var paragraphs []string
                // Find the full move body, assign to a created paragraphs slice.
                container.Find(".attack-info-body > p").Each(func(_ int, paragraph *goquery.Selection) {
                      text := visibleText(paragraph)
                      if text != "" {
                      	paragraphs = append(paragraphs, text)
                      }   
                })

                container.PrevAll().Filter("p").EachWithBreak(func(i int, p *goquery.Selection) bool {
                        fmt.Printf("previous p[%d]: %q\n", i, strings.TrimSpace(p.Text()))
                        return i < 5
                })

                // Join the strings in the created paragraphs slice, assign to move.Description
                move.Description = strings.Join(paragraphs, "\n")

                // Append the created struct to the Moves Slice
                moves = append(moves, move)
        })

        return moves, nil
}

// Helper to parse goquery data
func parseCell(cell *goquery.Selection) models.Cell {
        toolTip := strings.TrimSpace(cell.Find(".tooltiptext").Text())

        visible := cell.Clone()
        visible.Find(".tooltiptext").Remove()

        return models.Cell{
                Value:  strings.TrimSpace(visible.Text()),
                Tooltip: toolTip,
        }
}

// Helper to seperate HTML backend tooltips from visible text
func visibleText(selection *goquery.Selection) string {
	// Clone the input go query
	clone := selection.Clone()

	// Find each ".tooltip", clone and strip trailing text, replace text with stripped versions
	clone.Find(".tooltip").Each(func(_ int, tooltip *goquery.Selection) {
		label := strings.TrimSpace(
			tooltip.Children().Not(".tooltiptext").Text(),
		)
		tooltip.ReplaceWithHtml(label)
	})

	// Rejoin text fields
	return strings.Join(strings.Fields(clone.Text()), " ")
}
