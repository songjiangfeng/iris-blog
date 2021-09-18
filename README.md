# iris-blog


$  git clone https://github.com/songjiangfeng/iris-blog.git

## 安装 Iris CLI 
$ go get github.com/kataras/iris-cli


## 创建数据库

$ mysqladmin create iris_blog
## 导入数据库

$ mysql -uroot -p iris_blog < iris_blog.sql

## 配置数据库文件

$ cp config.json.default config.json

## 运行  http://localhost:8080
$ iris-cli run .


## 后台 http://localhost:8080/admin  （注意记得更改密码）
用户名： admin
密码： admin

