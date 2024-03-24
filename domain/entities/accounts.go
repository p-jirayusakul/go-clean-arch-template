package entities

type Accounts struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountsDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountsIdDto struct {
	ID string `json:"id"`
}
