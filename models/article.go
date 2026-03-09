package models

type Article struct {
	ID       int     `db:"id" json:"id"`
	Title    string  `db:"title" json:"title"`
	Slug     string  `db:"slug" json:"slug"`
	Content  string  `db:"content" json:"content"`
	AuthorID int     `db:"author_id" json:"author_id"`
	Status   *string `db:"status" json:"status"`
}

type ArticleWithAuthor struct {
	ID         int     `db:"id" json:"id"`
	Title      string  `db:"title" json:"title"`
	Slug       string  `db:"slug" json:"slug"`
	Content    string  `db:"content" json:"content"`
	AuthorName *string `db:"author_name" json:"author_name"`
}
