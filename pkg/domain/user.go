package domain

type User struct {
	Username             string       `bson:"username" json:"username"`
	FavouriteRecipes     []string     `bson:"favouriteRecipes" json:"favouriteRecipes"`
	AvailableIngredients []Ingredient `bson:"availableIngredients" json:"availableIngredients"`
}
