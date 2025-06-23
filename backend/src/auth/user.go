package auth

const connString = "postgres://user:password@localhost:5432/mydb"

type User struct {
	Id         int
	Email      string
	FirstName  string
	LastName   string
	PictureURL string
}
