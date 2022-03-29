# 博客系统后台

技术栈 `echo` + `gorm`

## 集成功能
- [x] JWT 验证
- [x] rbac角色权限认证
- [x] 用户注册
- [x] 用户登录
- [x] 接口文档
- [x] 用户分类增删改查
- [x] 用户分类文章增删改查
- [x] 用户评论增删改查
- [x] 日志记录

## docker 部署方式

```shell
docker build -t golang_blog:v1.0.0 .
```

```shell
docker run -d --rm -p 8080:8080 --name golang_blog_project -v E:\Desktop\golang_blog:/app/golang_blog  --link MySQL:mysqldb golang_blog:v1.0.0
```

## 访问接口文档
http://127.0.0.1:8080/swagger/index.html