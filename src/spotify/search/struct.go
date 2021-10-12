package search

type TrackResponse struct {
	Tracks struct {
		Items []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Popularity int    `json:"popularity"`
			Album      struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				ReleaseDate string `json:"release_date"`
			} `json:"album"`
			Artists []struct {
				ID string `json:"id"`
			} `json:"artists"`
		} `json:"items"`
	} `json:"tracks"`
}
