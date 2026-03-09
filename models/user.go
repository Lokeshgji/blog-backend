package models

type User struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	Role      string `db:"role" json:"role"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
