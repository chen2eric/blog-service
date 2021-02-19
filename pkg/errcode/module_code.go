package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签列表失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签列表失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签列表失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签列表失败")

	ErrorCreateArticleFail  = NewError(20020001, "创建文章失败")
	ErrorDeleteArticleFail  = NewError(20020002, "删除文章失败")
	ErrorUpdateArticleFail  = NewError(20020003, "更新文章失败")
	ErrorGetOneArticleFail  = NewError(20020004, "获取单个文章失败")
	ErrorGetArticleListFail = NewError(20020005, "获取文章列表失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")
)
