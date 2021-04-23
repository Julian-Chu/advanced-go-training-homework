package user

type User struct {
	ID           string         `db:"user_id" json:"id"`
}