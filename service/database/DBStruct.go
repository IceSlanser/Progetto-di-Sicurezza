package database

type User struct {
	ID       uint64 `json:"ID"`
	Username string `json:"username"`
	Password string `json:"password"`
}
