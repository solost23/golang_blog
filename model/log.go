package model

import "time"

type Log struct {
	ID            int32  `gorm:"column:id;primary_key"`
	UserName      string `gorm:"column:user_name"`
	OperationType string `gorm:"column:operation_type;type:enum('NONE','INSERT','DELETE','UPDATE');default:NONE"`
	SourceType    string `gorm:"column:source_type;type:enum('NONE','CONTENT','USER','ARTICLE','COMMENT');default:NONE'"`
	SourceName    string `gorm:"column:source_name"`
	IsSuccess     string `gorm:"column:is_success;type:enum('NONE','SUCCESS','FAILED');default:NONE"`
	CreateTime    int64  `gorm:"column:create_time"`
	UpdateTime    int64  `gorm:"column:update_time"`
}

func (l *Log) TableName() string {
	return "log"
}

// 创建一条日志
func (l *Log) Create() error {
	l.CreateTime = time.Now().Unix()
	l.UpdateTime = time.Now().Unix()
	if err := DB.Table("log").Create(l).Error; err != nil {
		return err
	}
	return nil
}

// 查询所有log
func (l *Log) Find() ([]*Log, error) {
	var res []*Log
	if err := DB.Table("log").Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// 根据id删除一条日志
func (l *Log) DeleteByID(ID int32) error {
	if err := DB.Table("log").Where("id=?", ID).Delete(l).Error; err != nil {
		return err
	}
	return nil
}

// 根据id查询本条日志是否存在
func (l *Log) FindByID(ID int32) error {
	if err := DB.Table("log").Where("id=?", ID).First(l).Error; err != nil {
		return err
	}
	return nil
}
