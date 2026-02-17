package user

type UpdateUserInput struct {
	ID    string
	Name  string
	Email string
}

type UpdateUserOutput struct {
	ID    string
	Name  string
	Email string
}
