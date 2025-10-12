package domain

type User struct {
	Id             int64
	Name           string
	Surname        string
	Email          string
	HashedPassword string
	Username       string
	Role           string
	CreatedAt      string
	UpdatedAt      string
}
