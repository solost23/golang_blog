package models

type Moder interface {
	TableName() string
	Insert(interface{}) error
	Delete(interface{}, ...interface{}) error
	Save(interface{}, interface{}, ...interface{}) error
	WhereOne(interface{}, ...interface{}) (interface{}, error)
	WhereAll(interface{}, ...interface{}) (interface{}, error)
	PageList(*ListPageInput, interface{}, ...interface{}) (interface{}, int64, error)
}