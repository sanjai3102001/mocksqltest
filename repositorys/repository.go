//
//  go-unit-test-sql
//
//

package repository

// Repository represent the repositories
type Repository interface {
	Close()
	FindByID(id string) (*UserModel, error)
	Find() ([]*UserModel, error)
	Create(user *UserModel) error
	Update(user *UserModel) error
	Delete(id string) error
}

// UserModel represent the user model
type UserModel struct {
	ID    string
	Name  string
	Email string
	Phone string
}
