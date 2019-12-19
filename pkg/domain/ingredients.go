package domain

type Ingredient struct {
	Name     string `bson:"ingredients" json:"ingredients"`
	Quantity int    `bson:"quantity" json:"quantity"`
}
