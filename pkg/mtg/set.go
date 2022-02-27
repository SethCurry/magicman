package mtg

type Set struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Border      string `json:"border"`
	ReleaseDate string `json:"releaseDate"`
}
