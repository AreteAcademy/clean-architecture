package domain

type Category struct {
	ID        string
	UserId    string
	Name      string
	Status    string
	CreatedAt string
	UpdatedAt string
}

type CategoryRepository interface {
	Save(category *Category) error
	Update(category *Category) error
	GetById(id string) (*Category, error)
	Count() (int, error)
}
