package domain

type UserRepository interface {
	Save(user User) error
	Get(userName string) (*User, error)
}
