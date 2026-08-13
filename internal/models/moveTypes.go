package models

type Move struct {
        Name            string
        Input           string
        Headers         []string
        FrameData       map[string]Cell
        Notes           []string
        Description     string
}

type Cell struct {
        Value           string
        Tooltip         string
}
