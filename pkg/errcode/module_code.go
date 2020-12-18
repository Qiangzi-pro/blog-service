package errcode

var (
	ErrorGetTagListFail = NewError(20001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20004, "删除标签失败")
	ErrorCountTagFail   = NewError(20005, "统计标签失败")

	ErrorGetArticleFail    = NewError(30001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(30002, "创建多个文章失败")
	ErrorCreateArticleFail = NewError(30003, "创建文件失败")
	ErrorUpdateArticleFail = NewError(30004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(30005, "删除文章失败")
)
