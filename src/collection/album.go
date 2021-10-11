package collection

type Album struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	ReleaseDate string `json:"release_date,omitempty" bson:"release_date,omitempty"`
}
