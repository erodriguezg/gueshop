package users

type User struct {
    ID    int64  `db:"id" json:"id"`
    Name  string `db:"name" json:"name"`
    Email string `db:"email" json:"email"`
}

