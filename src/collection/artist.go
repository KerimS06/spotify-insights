package collection

type Artist struct {
	ID         string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string   `json:"name,omitempty" bson:"name,omitempty"`
	Popularity int      `json:"popularity,omitempty" bson:"popularity,omitempty"`
	Genres     []string `json:"genres,omitempty" bson:"genres,omitempty"`
	Followers  int      `json:"followers,omitempty" bson:"followers,omitempty"`
}
