package workList

const (
	NONE    = "NONE"
	INSERT  = "INSERT"
	DELETE  = "DELETE"
	UPDATE  = "UPDATE"
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
	CONTENT = "CONTENT"
	USER    = "USER"
	ARTICLE = "ARTICLE"
	COMMENT = "COMMENT"
)

type WorkList struct{}

func NewWorkList() *WorkList {
	return &WorkList{}
}
