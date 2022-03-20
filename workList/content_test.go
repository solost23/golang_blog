package workList

import (
	"log"
	"testing"

	"golang_blog/model"
)

func TestWorkList_GetAllContent(t *testing.T) {
	var workList *WorkList
	var content *model.Content
	contentList, err := workList.GetAllContent(content)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(contentList)
}
