package model

type DBInterface interface {
	Create() error
	Delete() error
	Update() error
	Find() ([]*Content, error)
	FindByName(string) error
	FindById() error
}
