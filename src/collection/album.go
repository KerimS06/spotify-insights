package collection

type Album struct {
	ID          string `json:"id,omitempty" bson:"id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	ReleaseDate string `json:"release_date,omitempty" bson:"release_date,omitempty"`
}
