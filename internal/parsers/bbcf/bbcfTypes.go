package bbcf

// Define Structs to handle Character data
type BBCFCharData struct {
	Name		string
	Overview	string
	Mechanics	[]Mechanic
	Moves		[]Move
	Infobox		map[string]string
}

type Mechanic struct {
	Name		string
	Description	string
}

type Move struct {
	Name		string
	Headers		[]string
	FrameData	map[string]Cell
	Notes		[]string
	Description	string
}

type Cell struct {
	Value		string
	Tooltip		string
}
