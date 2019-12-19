package domain

type Repository interface {
	FindUser(id string) (*User, error)
	CreateUser(user User) error
	UpdateUser(user User) error

	FindRecipes(restrictions Restrictions) ([]*Recipe, error)
}
