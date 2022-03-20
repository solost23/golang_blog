package workList

import (
	"testing"
)

func TestWorkList_CreateLog(t *testing.T) {
	var w *WorkList
	if err := w.CreateLog("ty", INSERT, ARTICLE, "golang", SUCCESS); err != nil {
		t.Log(err.Error())
	}
	t.Log("success")
}
