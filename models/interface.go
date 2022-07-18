package models

type Moder interface {
	TableName() string
	Insert() error
	Delete(interface{}, ...interface{}) error
	Save(interface{}, ...interface{}) error
	WhereOne(interface{}, ...interface{}) (interface{}, error)
	WhereAll(interface{}, ...interface{}) (interface{}, error)
	PageList(*ListPageInput, interface{}, ...interface{}) (interface{}, int64, error)
}
