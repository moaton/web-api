package user

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}
