package model

type Moder interface {
	TableName() string
	Insert(interface{}) error
	Delete(interface{}, ...interface{}) error
	Save(data interface{}) error
	WhereOne(interface{}, ...interface{}) (interface{}, error)
	WhereAll(interface{}, ...interface{}) (interface{}, error)
	PageList(*ListPageInput, interface{}, ...interface{}) (interface{}, int64, error)
}
