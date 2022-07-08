package workList

import (
	"log"
	"testing"

	"golang_blog/models"
)

func TestWorkList_GetAllContent(t *testing.T) {
	var workList *WorkList
	var content *models.Content
	contentList, err := workList.GetAllContent(content)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(contentList)
}
