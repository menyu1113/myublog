# myublog

### 介绍
博客后端接口
### 技术栈

### 软件架构
软件架构说明
```text
├── api         
├─config // 项目配置入口   
├─database  // 数据库备份文件（初始化）
├─global  //全局配置
│  ├─myerrors 自定义错误及其他常量
│  ├─gorm //mysql初始化
│  ├─initredis //redis初始化
│  ├─viperis 配置加载
│  └─global // 全局变量声明
├─logs  // 日志路径
├─middleware  // 中间件
│  ├─loglog //日志中间件
│  ├─cors //跨域中间件
│  └─jwt //token中间件
├─model // 模型层
├─routes
│  └─router.go // 路由入口   
├─service // 服务层             
├─utils // 工具库
│  ├─md5x //md5加密 
│  ├─response //数据返回响应   
│  ├─validators //参数验证
│  ├─zaplog //日志工具      
│  └─utiljwt //jwt 
├─ air.conf  
├─  .gitignore
├─ go.mod // 项目依赖
├─ go.sum
├─ main.go //主程序
├─ README.md 
```
### 表设计
####用户表->博客表 一对一
####用户表->用户详情表 一对一
####用户表->分类表 一对多
####用户表->评论表 一对多
####用户表->文章表 一对多
####分类表->文章表 一对多
####文章表->评论表 一对多

### 使用说明
##### 1、配置config下面的config.yaml，注意的是日志如果更改路径需要自己创建
##### 2、cd myublog
##### 3、go run main.go
