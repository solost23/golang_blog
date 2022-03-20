package workList

import (
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"golang_blog/model"
)

func (w *WorkList) GetAllLog(log *model.Log) ([]*model.Log, error) {
	// 直接获取日志并返回
	logList, err := log.Find()
	if err != nil {
		return logList, err
	}
	return logList, err
}

func (w *WorkList) DeleteLog(log *model.Log) error {
	// 查询日志是否存在，若不存在，则返回错误
	// 若存在，则删除
	id := w.ctx.Get("id").(string)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := log.FindByID(int32(idInt)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(err.Error())
			return err
		}
	}
	if err := log.DeleteByID(int32(idInt)); err != nil {
		return err
	}
	return nil
}

// 调用log,存入数据库
func (w *WorkList) CreateLog(userName, operationType, sourceType, sourceName, isSuccess string) error {
	log := &model.Log{
		UserName:      userName,
		OperationType: operationType,
		SourceType:    sourceType,
		SourceName:    sourceName,
		IsSuccess:     isSuccess,
	}
	if err := log.Create(); err != nil {
		return err
	}
	return nil
}
