package repository

type Authorization interface {
}

type AppList interface {
}

type AppItem interface {
}

type Repository struct {
	Authorization
	AppList
	AppItem
}

func NewRepository() *Repository {
	return &Repository{}
}
