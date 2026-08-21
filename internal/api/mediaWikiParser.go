package api


type MediaWikiParseResponse struct {
	Parse struct {
		Title string `json:"title"`
		Text struct {
			HTML string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
	Error *struct {
		Code string `json:"code"`
		Info string `json:"info"`
	} `json:"error,omitempty"`
}
