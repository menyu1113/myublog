package consts

const (
	SUCCSECODE int    = 200
	SUCCSE     string = "ok"
	ERRORCODE  int    = 201
	ERROR      string = "FAIL"

	CurdCreatFailExistCode int    = 1001
	CurdCreatFailExist     string = "用户名已存在"

	CurdCreatFailCode int    = 1002
	CurdCreatFailMsg  string = "新增失败"

	ERROR_PASSWORD_Code int    = 1003
	ERROR_PASSWORD_Msg  string = "密码错误"

	PasswordChangeFailed    string = "密码修改失败"
	CurdUpdateFailCode      int    = 1004
	CurdUpdateFailMsg       string = "更新失败"
	CurdUpdateFailExistCode int    = 1005
	CurdUpdateFailExistMsg  string = "更新失败,用户名已存在"

	CurdDeleteFailCode       int    = 1006
	CurdDeleteFailMsg        string = "删除失败"
	CurdSelectFailCode       int    = 1007
	CurdSelectFailMsg        string = "查询无数据"
	CurdRegisterFailCode     int    = 1008
	CurdRegisterFailMsg      string = "注册失败"
	CurdLoginFailCode        int    = 1009
	CurdLoginFailMsg         string = "登录失败"
	CurdRefreshTokenFailCode int    = 1010
	CurdRefreshTokenFailMsg  string = "刷新Token失败"
	BindParameterCode        int    = 1011
	BindParameter            string = "参数绑定失败"
	CurdFailNotExistCode     int    = 1012
	CurdFailNotExist         string = "用户ID不存在"
	LoginUserNotExistCode    int    = 1013
	LoginUserNotExist        string = "用户不存在"
	FailedAddArticleCode     int    = 1014
	FailedAddArticle         string = "添加文章失败"

	FailGetUserDetailsCode    int    = 1015
	FailGetUserDetails        string = "获取用户详情失败"
	FailUpdataUserDetailsCode int    = 1016
	FailUpdataUserDetails     string = "更新用户详情失败"
	FailUserLinkDetailsCode   int    = 1017
	FailUserLinkDetails       string = "关联失败"
	FailGetArticleCode        int    = 1018
	FailGetArticle            string = "获取文章失败"
	FailUpdataArticleCode     int    = 1019
	FailUpdataArticle         string = "编辑文章失败"
	FailDeleArticleCode       int    = 1020
	FailDeleArticle           string = "删除文章失败"
	FailGetCategoryCode       int    = 1021
	FailGetCategory           string = "获取分类列表失败"
	FailCatenameUsedCode      int    = 1022
	FailCatenameUsed          string = "分类已存在"
	FailUpdataCateCode        int    = 1023
	FailUpdataCate            string = "编辑分类失败"
	FailDeleCateCode          int    = 1024
	FailDeleCate              string = "删除分类失败"
	FailGetCommentCode        int    = 1025
	FailGetComment            string = "获取评论失败"
	FailDeleCommentCode       int    = 1026
	FailDeleComment           string = "删除评论失败"
)

var codeMsg = map[int]string{
	SUCCSECODE:                "OK",
	ERRORCODE:                 "FAIL",
	CurdCreatFailExistCode:    "用户名已存在",
	CurdUpdateFailExistCode:   "更新失败,用户名已存在",
	CurdCreatFailCode:         "新增失败",
	ERROR_PASSWORD_Code:       "密码错误",
	CurdUpdateFailCode:        "更新失败",
	CurdDeleteFailCode:        "删除失败",
	CurdSelectFailCode:        "查询无数据",
	CurdLoginFailCode:         "登录失败",
	BindParameterCode:         "参数绑定失败",
	CurdFailNotExistCode:      "用户ID不存在",
	LoginUserNotExistCode:     "用户不存在",
	FailedAddArticleCode:      "添加文章失败",
	FailGetUserDetailsCode:    "获取用户详情失败",
	FailUpdataUserDetailsCode: "更新用户详情失败",
	FailUserLinkDetailsCode:   "关联失败",
	FailGetArticleCode:        "获取文章失败",
	FailUpdataArticleCode:     "编辑文章失败",
	FailDeleArticleCode:       "删除文章失败",
	FailGetCategoryCode:       "获取分类列表失败",
	FailCatenameUsedCode:      "分类已存在",
	FailUpdataCateCode:        "编辑分类失败",
	FailDeleCateCode:          "删除分类失败",
	FailGetCommentCode:        "获取评论失败",
	FailDeleCommentCode:       "删除评论失败",
	//ERROR_USER_NOT_EXIST:   "用户不存在",
	//ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	//ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	//ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	//ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	//ERROR_USER_NO_RIGHT:    "该用户无权限",
	//
	//ERROR_ART_NOT_EXIST: "文章不存在",
	//
	//ERROR_CATENAME_USED:  "该分类已存在",
	//ERROR_CATE_NOT_EXIST: "该分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
