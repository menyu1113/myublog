package myerrors

const (
	SUCCSECODE int    = 200
	SUCCSE     string = "ok"
	ERRORCODE  int    = 201
	ERROR      string = "FAIL"


	ErrorsGormInitFail             string = "Gorm 数据库驱动、连接初始化失败"
	ErrorsGormNotInitGlobalPointer string = "%s 数据库全局变量指针没有初始化，请在配置文件 Gormv2.yml 设置 Gormv2.%s.IsInitGolobalGormMysql = 1, 并且保证数据库配置正确 \n"
	// 数据库部分
	ErrorsDbDriverNotExists        string = "数据库驱动类型不存在,目前支持的数据库类型：mysql、sqlserver、postgresql，您提交数据库类型："
	ErrorsDialectorDbInitFail      string = "gorm dialector 初始化失败,dbType:"
	ErrorsGormDBCreateParamsNotPtr string = "gorm Create 函数的参数必须是一个指针"
	ErrorsGormDBUpdateParamsNotPtr string = "gorm 的 Update、Save 函数的参数必须是一个指针(GinSkeleton ≥ v1.5.29 版本新增验证，为了完美支持 gorm 的所有回调函数,请在参数前面添加 & )"
	//redis部分
	ErrorsRedisInitConnFail string = "初始化redis连接池失败"
	ErrorsRedisAuthFail     string = "Redis Auth 鉴权失败，密码错误"
	ErrorsRedisGetConnFail  string = "Redis 从连接池获取一个连接失败，超过最大重试次数"
	// 表单参数验证器未通过时的错误
	ErrorsValidatorNotExists      string = "不存在的验证器"
	ErrorsValidatorTransInitFail  string = "validator的翻译器初始化错误"
	ErrorNotAllParamsIsBlank      string = "该接口不允许所有参数都为空,请按照接口要求提交必填参数"
	ErrorsValidatorBindParamsFail string = "验证器绑定参数失败"

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
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
