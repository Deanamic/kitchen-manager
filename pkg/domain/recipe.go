package domain

type Recipe struct {
	Name         string       `bson:"name" json:"name"`
	Ingredients  []Ingredient `bson:"ingredients" json:"ingredients"`
	Instructions []string     `bson:"instructions" json:"instructions"`
}
