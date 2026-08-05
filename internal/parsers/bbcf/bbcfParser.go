package bbcf

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func BBCFCharPageParser(data []byte) ([]Move, error) {

	// Use goquery to create a readable datapoint from page HTML data
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Initialise a slice of Move structs
	var moves []Move
	
	headings := doc.Find("h3")
	// Find all attack containers in the created doc
	doc.Find(".attack-container").Each(func(i int, container *goquery.Selection) {

		// Initialise an internal headers slice for matching data points
		headers := []string{}

		// Find all of the framedatagridheaders in the doc
		container.Find(".frameDataGridHeader > div").Each(func(_ int, cell *goquery.Selection) {
			// Clone the created goquery
			visible := cell.Clone()

			// Find the tooltip text inside the copied goquery and remove it
			visible.Find(".tooltiptext").Remove()

			// Append the remaining text (the raw header without explanation) to the headers slice
			headers = append(headers, strings.TrimSpace(visible.Text()))
		})

		// Initialise the Framedata map for a Move struct
		move := Move{
			FrameData: make(map[string]Cell),
			Headers: []string{},
		}

		// Find and assign the move name to the move.Name element
		if i < headings.Length() {
			heading := headings.Eq(i)
			move.Name = strings.TrimSpace(heading.Text())		
		}		

		// Find the framedata grid rows present in the Movecontainers
		container.Find(".frameDataGridRow > div").Each(func(i int, cell *goquery.Selection) {

			// if headers is shorter than the number of framedata rows, safetly return 
			if i >= len(headers) {
				return
			}

			// Parse the created goquery with the helper and assign the header data to move.FrameData
			move.FrameData[move.Headers[i]] = parseCell(cell)
		})

		// Find the attack body info containing the move tooltips from the wiki
		container.Find(".attack-info-body > ul > li").Each(func(_ int, note *goquery.Selection) {
			// Append the text element of the created goquery to the Notes element of the created Move struct
			move.Notes = append(move.Notes, strings.TrimSpace(note.Text()))
		})

		var paragraphs []string
		// Find the full move body, assign to a created paragraphs slice. 
		container.Find(".attack-info-body > p").Each(func(_ int, paragraph *goquery.Selection) {
			text := strings.TrimSpace(paragraph.Text())
			if text != "" {
				paragraphs = append(paragraphs, text)
			}
		})
		
		// Join the strings in the created paragraphs slice, assign to move.Description
		move.Description = strings.Join(paragraphs, "\n")

		// Append the created struct to the Moves Slice
		moves = append(moves, move)
	})

	return moves, nil
}

// Helper to parse goquery data
func parseCell(cell *goquery.Selection) Cell {
	toolTip := strings.TrimSpace(cell.Find(".tooltiptext").Text())

	visible := cell.Clone()
	visible.Find(".tooltiptext").Remove()

	return Cell{
		Value:	strings.TrimSpace(visible.Text()),
		Tooltip: toolTip,
	}
}
