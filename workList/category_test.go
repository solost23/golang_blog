package workList

import (
	"log"
	"testing"

	"golang_blog/models"
)

func TestWorkList_GetAllCategory(t *testing.T) {
	var workList *WorkList
	var category *models.Category
	contentList, err := workList.GetAllCategory(category)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(contentList)
}
