# gin 写的简单网站
用于展示爬取的节点

[节点分享链接](https://github.com/sanzhang007/node_free)
## models 数据库
### core 连接数据库
// 修改为你的配置
```go
cmd := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local" 
//改为自己的配置 root:123456对应用户和密码 test对应数据库名
//sql语句创建 数据库 create database database_name
```
### nodes 表结构

### protocol 代理协议 结构体

### template clash 模板

### controlers 模块放(gin handlers)

