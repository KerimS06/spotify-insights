package collection

type Track struct {
	ID         string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string   `json:"name,omitempty" bson:"name,omitempty"`
	Popularity int      `json:"popularity,omitempty" bson:"popularity,omitempty"`
	Album      Album    `json:"album,omitempty" bson:"album,omitempty"`
	Artists    []Artist `json:"artists,omitempty" bson:"artists,omitempty"`
}
