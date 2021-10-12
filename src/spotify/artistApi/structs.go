package artistApi

type ArtistResponse struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Popularity int      `json:"popularity"`
	Genres     []string `json:"genres"`
	Followers  struct {
		Total int `json:"total"`
	} `json:"followers"`
}
