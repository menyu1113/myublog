## /myublog
```text
后端API
```
## /myublog/用户
```text
暂无描述
```
#### 公共Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/用户/注册用户详情
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/adddetail/1

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
	"desc": "描述信息",
	"wechat": "qp123",
	"email": "123qq.com",
	"avatar":"头像"
}
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/用户/删除用户
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/user/6

#### 请求方式
> DELETE

#### Content-Type
> json

#### 请求Body参数
```javascript

```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
username | qazwsx02 | String | 是 | 用户名,4-12位
password | 123456 | String | 是 | 密码,4-128位
role | 3 | Number | 是 | 角色属性，必须大于等于2
age | 20 | Text | 是 | -
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "OK"
}
```
## /myublog/用户/更改用户信息
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/user/2

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Body参数
```javascript
{
	"username": "qazwsx20",
	"password": "123456",
	"role": 3,
	"age":20
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
username | qazwsx02 | String | 是 | 用户名,4-12位
password | 123456 | String | 是 | 密码,4-128位
role | 3 | Number | 是 | 角色属性，必须大于等于2
age | 20 | Text | 是 | -
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/用户/用户注册
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/user/add

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
	"username": "qazwsx02",
	"password": "123456",
	"role": 2,
	"age":20
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
username | qazwsx02 | String | 是 | 用户名,4-12位
password | 123456 | String | 是 | 密码,4-128位
role | 3 | Number | 是 | 角色属性，必须大于等于2
age | 20 | Text | 是 | -
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/用户/更改用户密码
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/user/changepasswd/2

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Body参数
```javascript
{
	"username": "qazwsx01",
	"password": "1234567",
	"role": 3,
	"age":20
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
username | qazwsx02 | String | 是 | 用户名,4-12位
password | 123456 | String | 是 | 密码,4-128位
role | 3 | Number | 是 | 角色属性，必须大于等于2
age | 20 | Text | 是 | -
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/用户/查询用户列表
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/admin/users?pagesize=5&pagenum=2

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
pagesize | 5 | Text | 是 | -
pagenum | 2 | Text | 是 | -
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/用户/登录
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/login

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
	"username": "qazwsx01",
	"password": "123456",
	"role": 3,
	"age":20
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
username | qazwsx02 | String | 是 | 用户名,4-12位
password | 123456 | String | 是 | 密码,4-128位
role | 3 | Number | 是 | 角色属性，必须大于等于2
age | 20 | Text | 是 | -
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/用户/获取用户详情
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/admin/detail/1

#### 请求方式
> GET

#### Content-Type
> json

#### 请求Body参数
```javascript

```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		[
			{
				"ID": 1,
				"CreatedAt": "2022-04-13T02:15:50.064+08:00",
				"UpdatedAt": "2022-04-13T02:15:50.187+08:00",
				"DeletedAt": null,
				"UserID": 1,
				"User": {
					"ID": 1,
					"CreatedAt": "2022-04-13T02:15:10.962+08:00",
					"UpdatedAt": "2022-04-13T02:15:10.962+08:00",
					"DeletedAt": null,
					"username": "qazwsx01",
					"password": "$2a$10$gAxIifpOOv0tudIGpQ9Or.36N/QvqwPiqyEmoepL.KmiU0f1mbTn6",
					"role": 2
				},
				"desc": "描述信息",
				"wechat": "qp123",
				"email": "123qq.com",
				"img": "",
				"avatar": "头像"
			}
		]
	],
	"msg": "OK"
}
```
## /myublog/分类
```text
暂无描述
```
#### 公共Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/分类/添加分类
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/category/add/1

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
	"name":"分类名称4"
}
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": null,
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "关联失败"
}
```
## /myublog/分类/查看分类列表
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/admin/category?pagesize=1&pagenum=1

#### 请求方式
> GET

#### Content-Type
> json

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
pagesize | 1 | Text | 是 | -
pagenum | 1 | Text | 是 | -
#### 请求Body参数
```javascript

```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		[
			[
				{
					"Articles": null,
					"UserID": 1,
					"id": 1,
					"name": "分类名称"
				},
				{
					"Articles": null,
					"UserID": 1,
					"id": 2,
					"name": "分类名称2"
				},
				{
					"Articles": null,
					"UserID": 1,
					"id": 3,
					"name": "分类名称3"
				},
				{
					"Articles": null,
					"UserID": 1,
					"id": 4,
					"name": "分类名称4"
				}
			],
			4
		]
	],
	"msg": "OK"
}
```
## /myublog/分类/编辑分类
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/category/2

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Body参数
```javascript
{
    "name":"修改后的分类名称"
}
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 1023,
	"data": null,
	"msg": "编辑分类失败"
}
{
	"code": 1022,
	"data": null,
	"msg": "分类已存在"
}
```
## /myublog/分类/删除分类
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/category/1?cateid=1

#### 请求方式
> DELETE

#### Content-Type
> json

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
cateid | 1 | Text | 是 | 分类id
#### 请求Body参数
```javascript
{
    "name":"修改后的分类名称"
}
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": null,
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 1024,
	"data": null,
	"msg": "删除分类失败"
}
```
## /myublog/文章
```text
暂无描述
```
#### 公共Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/文章/添加文章
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/article/add/1

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
	"title":"这是标题3",
	"desc":"这是描述3",
	"content":"这是具体文章内容3"
}
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "关联失败"
}
```
## /myublog/文章/获取用户所有文章
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/article/1

#### 请求方式
> GET

#### Content-Type
> json

#### 请求Body参数
```javascript

```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		[
			[
				{
					"ID": 1,
					"CreatedAt": "2022-04-14T02:05:30.756+08:00",
					"UpdatedAt": "2022-04-14T02:05:30.756+08:00",
					"DeletedAt": null,
					"CategoryID": 0,
					"UserID": 1,
					"Comments": null,
					"title": "这是标题",
					"desc": "这是描述",
					"content": "这是具体文章内容",
					"img": "",
					"comment_count": 0,
					"read_count": 0
				},
				{
					"ID": 2,
					"CreatedAt": "2022-04-14T02:51:45.257+08:00",
					"UpdatedAt": "2022-04-14T03:41:10.02+08:00",
					"DeletedAt": null,
					"CategoryID": 0,
					"UserID": 1,
					"Comments": null,
					"title": "编辑更改了标题",
					"desc": "编辑更改了描述",
					"content": "GORM 允许通过 Select 方法选择特定的字段，如果您在应用程序中经常使用此功能，你可以定义一个较小的 API 结构体，以实现自动选择特定的字段。",
					"img": "",
					"comment_count": 0,
					"read_count": 0
				},
				{
					"ID": 3,
					"CreatedAt": "2022-04-14T03:46:16.161+08:00",
					"UpdatedAt": "2022-04-14T03:46:16.161+08:00",
					"DeletedAt": null,
					"CategoryID": 0,
					"UserID": 1,
					"Comments": null,
					"title": "这是标题3",
					"desc": "这是描述3",
					"content": "这是具体文章内容3",
					"img": "",
					"comment_count": 0,
					"read_count": 0
				}
			]
		]
	],
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "关联失败"
}
```
## /myublog/文章/获取用户单个文章
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/article/info/1/?artid=3

#### 请求方式
> GET

#### Content-Type
> json

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
artid | 3 | Number | 是 | 文章id
#### 请求Body参数
```javascript

```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		[
			{
				"ID": 2,
				"CreatedAt": "2022-04-14T02:51:45.257+08:00",
				"UpdatedAt": "2022-04-14T02:51:45.257+08:00",
				"DeletedAt": null,
				"CategoryID": 0,
				"UserID": 1,
				"Comments": null,
				"title": "这是标题2",
				"desc": "这是描述2",
				"content": "这是具体文章内容2",
				"img": "",
				"comment_count": 0,
				"read_count": 0
			}
		]
	],
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "关联失败"
}
```
## /myublog/文章/编辑文章
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/article/1/2

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Body参数
```javascript
{
    "title":"编辑更改了标题",
    "desc":"编辑更改了描述",
    "content":"GORM 允许通过 Select 方法选择特定的字段，如果您在应用程序中经常使用此功能，你可以定义一个较小的 API 结构体，以实现自动选择特定的字段。"
}
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		[
			{
				"ID": 2,
				"CreatedAt": "2022-04-14T02:51:45.257+08:00",
				"UpdatedAt": "2022-04-14T02:51:45.257+08:00",
				"DeletedAt": null,
				"CategoryID": 0,
				"UserID": 1,
				"Comments": null,
				"title": "这是标题2",
				"desc": "这是描述2",
				"content": "这是具体文章内容2",
				"img": "",
				"comment_count": 0,
				"read_count": 0
			}
		]
	],
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "关联失败"
}
```
## /myublog/文章/删除文章
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/article/3

#### 请求方式
> DELETE

#### Content-Type
> json

#### 请求Body参数
```javascript

```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		[
			{
				"ID": 2,
				"CreatedAt": "2022-04-14T02:51:45.257+08:00",
				"UpdatedAt": "2022-04-14T02:51:45.257+08:00",
				"DeletedAt": null,
				"CategoryID": 0,
				"UserID": 1,
				"Comments": null,
				"title": "这是标题2",
				"desc": "这是描述2",
				"content": "这是具体文章内容2",
				"img": "",
				"comment_count": 0,
				"read_count": 0
			}
		]
	],
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "关联失败"
}
```
## /myublog/评论
```text
暂无描述
```
#### 公共Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /myublog/评论/添加评论
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/addcomment/1?articleid=3

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
articleid | 3 | Text | 是 | 文章id
#### 请求Body参数
```javascript
{
    "content":"第2条评论"
}
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "OK"
}
```
#### 错误响应示例
```javascript
{
	"code": 200,
	"data": [
		null
	],
	"msg": "FAIL"
}
```
## /myublog/评论/删除评论
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/delcomment/2

#### 请求方式
> DELETE

#### Content-Type
> form-data

#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": null,
	"msg": "OK"
}
```
## /myublog/评论/获取评论
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> 127.0.0.1:9000/api/v1/comment/?artid=5

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
artid | 5 | Text | 是 | -
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"data": [
		[
			{
				"ArticleID": 1,
				"UserID": 1,
				"ID": 1,
				"CreatedAt": "2022-04-15T03:08:24.023+08:00",
				"UpdatedAt": "2022-04-15T03:08:24.023+08:00",
				"DeletedAt": null,
				"user_id": 0,
				"article_id": 0,
				"article_title": "",
				"username": "",
				"content": "第一条评论",
				"status": 2
			},
			{
				"ArticleID": 1,
				"UserID": 1,
				"ID": 2,
				"CreatedAt": "2022-04-15T03:08:40.228+08:00",
				"UpdatedAt": "2022-04-15T03:08:40.228+08:00",
				"DeletedAt": null,
				"user_id": 0,
				"article_id": 0,
				"article_title": "",
				"username": "",
				"content": "第2条评论",
				"status": 2
			}
		]
	],
	"msg": "ok"
}
```
#### 错误响应示例
```javascript
{
	"code": 1025,
	"data": null,
	"msg": "获取评论失败"
}
```